package cobracurl

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// BuildClient creates an http.Client configured from cobra command flags.
func BuildClient(cmd *cobra.Command) (*http.Client, error) {
	transport := &http.Transport{}

	var tlsConfig *tls.Config

	if insecure, _ := cmd.Flags().GetBool("insecure"); insecure {
		tlsConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
	}

	if cacertFile, _ := cmd.Flags().GetString("cacert"); cacertFile != "" {
		caCert, err := os.ReadFile(cacertFile)
		if err != nil {
			return nil, fmt.Errorf("reading cacert: %w", err)
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to parse CA certificate from %s", cacertFile)
		}
		if tlsConfig == nil {
			tlsConfig = &tls.Config{} //nolint:gosec
		}
		tlsConfig.RootCAs = caCertPool
	}

	if certFile, _ := cmd.Flags().GetString("cert"); certFile != "" {
		keyFile, _ := cmd.Flags().GetString("key")
		if keyFile == "" {
			keyFile = certFile
		}
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, fmt.Errorf("loading client certificate: %w", err)
		}
		if tlsConfig == nil {
			tlsConfig = &tls.Config{} //nolint:gosec
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	if tlsConfig != nil {
		transport.TLSClientConfig = tlsConfig
	}

	// tcp-nodelay is already enabled by default in Go (since Go 1.5).
	noKeepalive, _ := cmd.Flags().GetBool("no-keepalive")
	keepaliveTime, _ := cmd.Flags().GetInt("keepalive-time")
	connectTimeout, _ := cmd.Flags().GetFloat64("connect-timeout")

	dialer := &net.Dialer{}
	if connectTimeout > 0 {
		dialer.Timeout = time.Duration(connectTimeout * float64(time.Second))
	}
	if !noKeepalive && keepaliveTime > 0 {
		dialer.KeepAlive = time.Duration(keepaliveTime) * time.Second
	}
	transport.DialContext = dialer.DialContext

	if noKeepalive {
		transport.DisableKeepAlives = true
	}

	if proxyStr, _ := cmd.Flags().GetString("proxy"); proxyStr != "" {
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	client := &http.Client{Transport: transport}

	if maxTime, _ := cmd.Flags().GetFloat64("max-time"); maxTime > 0 {
		client.Timeout = time.Duration(maxTime * float64(time.Second))
	}

	location, _ := cmd.Flags().GetBool("location")
	if !location {
		client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}
	} else {
		maxRedirs, _ := cmd.Flags().GetInt("max-redirs")
		if maxRedirs > 0 {
			client.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
				if len(via) >= maxRedirs {
					return http.ErrUseLastResponse
				}
				return nil
			}
		}
	}

	return client, nil
}
