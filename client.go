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

// ClientOption customizes the client built by BuildClient.
type ClientOption func(*clientOptions)

type clientOptions struct {
	defaultTimeout time.Duration
}

// WithDefaultTimeout sets an overall request timeout used when --max-time is
// not set. The --max-time flag always takes precedence. curl has no overall
// timeout by default, so BuildClient applies none unless this option is given.
func WithDefaultTimeout(d time.Duration) ClientOption {
	return func(o *clientOptions) { o.defaultTimeout = d }
}

// BuildClient creates an http.Client configured from cobra command flags.
//
// The transport is cloned from http.DefaultTransport so it keeps the standard
// library defaults (proxy from environment, TLS handshake and idle timeouts,
// connection pooling). When --proxy is unset, HTTP_PROXY, HTTPS_PROXY and
// NO_PROXY are honored.
func BuildClient(cmd *cobra.Command, opts ...ClientOption) (*http.Client, error) {
	var options clientOptions
	for _, opt := range opts {
		opt(&options)
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()

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
		// A custom TLS config disables HTTP/2 unless explicitly re-enabled.
		transport.ForceAttemptHTTP2 = true
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

	if options.defaultTimeout > 0 {
		client.Timeout = options.defaultTimeout
	}
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
