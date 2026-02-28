package cobracurl

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

// BuildClient creates an http.Client configured from cobra command flags.
// Unlike the default http.Client, redirects are NOT followed unless --location
// is set, matching curl's default behavior.
func BuildClient(cmd *cobra.Command) (*http.Client, error) {
	transport := &http.Transport{}

	if insecure, _ := cmd.Flags().GetBool("insecure"); insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
	}

	if connectTimeout, _ := cmd.Flags().GetFloat64("connect-timeout"); connectTimeout > 0 {
		d := time.Duration(connectTimeout * float64(time.Second))
		transport.DialContext = (&net.Dialer{Timeout: d}).DialContext
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
