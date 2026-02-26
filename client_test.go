package cobracurl

import (
	"net/http"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildClient(t *testing.T) {
	tests := []struct {
		name             string
		flags            map[string]interface{}
		assertFn         func(t *testing.T, client *http.Client)
		expectedErrorMsg string
	}{
		{
			name:  "Default client does not follow redirects",
			flags: map[string]interface{}{},
			assertFn: func(t *testing.T, client *http.Client) {
				assert.Equal(t, time.Duration(0), client.Timeout)
				require.NotNil(t, client.CheckRedirect)
				assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, nil))
			},
		},
		{
			name: "Insecure skips TLS verification",
			flags: map[string]interface{}{
				"insecure": true,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				require.NotNil(t, transport.TLSClientConfig)
				assert.True(t, transport.TLSClientConfig.InsecureSkipVerify)
			},
		},
		{
			name: "Location enables redirect following",
			flags: map[string]interface{}{
				"location": true,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				assert.Nil(t, client.CheckRedirect)
			},
		},
		{
			name: "Location with max-redirs limits redirects",
			flags: map[string]interface{}{
				"location":   true,
				"max-redirs": 3,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				require.NotNil(t, client.CheckRedirect)
				via := make([]*http.Request, 3)
				assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, via))
				via = make([]*http.Request, 2)
				assert.NoError(t, client.CheckRedirect(nil, via))
			},
		},
		{
			name: "Max-time sets client timeout",
			flags: map[string]interface{}{
				"max-time": 30.0,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				assert.Equal(t, 30*time.Second, client.Timeout)
			},
		},
		{
			name: "Fractional max-time is converted correctly",
			flags: map[string]interface{}{
				"max-time": 1.5,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				assert.Equal(t, 1500*time.Millisecond, client.Timeout)
			},
		},
		{
			name: "Proxy sets transport proxy",
			flags: map[string]interface{}{
				"proxy": "http://proxy.example.com:8080",
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.NotNil(t, transport.Proxy)
			},
		},
		{
			name: "Invalid proxy URL returns error",
			flags: map[string]interface{}{
				"proxy": "://invalid",
			},
			expectedErrorMsg: "missing protocol scheme",
		},
		{
			name: "Connect-timeout sets dial timeout",
			flags: map[string]interface{}{
				"connect-timeout": 5.0,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.NotNil(t, transport.DialContext)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for key, value := range tt.flags {
				switch v := value.(type) {
				case bool:
					cmd.Flags().Bool(key, v, "")
				case string:
					cmd.Flags().String(key, v, "")
				case float64:
					cmd.Flags().Float64(key, v, "")
				case int:
					cmd.Flags().Int(key, v, "")
				}
			}

			client, err := BuildClient(cmd)

			if tt.expectedErrorMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, client)
			tt.assertFn(t, client)
		})
	}
}
