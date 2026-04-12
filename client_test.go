package cobracurl

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateTestCert creates a self-signed certificate and returns the PEM-encoded
// cert and key bytes. The certificate acts as both CA and client cert.
func generateTestCert(t *testing.T) (certPEM, keyPEM []byte) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		IsCA:         true,
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	require.NoError(t, err)
	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	keyDER, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})

	return certPEM, keyPEM
}

func writeTempFile(t *testing.T, data []byte) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "cobracurl-test-*")
	require.NoError(t, err)
	_, err = f.Write(data)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	return f.Name()
}

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
		{
			name: "no-keepalive disables keep-alives",
			flags: map[string]interface{}{
				"no-keepalive": true,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.True(t, transport.DisableKeepAlives)
			},
		},
		{
			name: "keepalive-time sets dialer keepalive interval",
			flags: map[string]interface{}{
				"no-keepalive":   false,
				"keepalive-time": 30,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.NotNil(t, transport.DialContext)
				assert.False(t, transport.DisableKeepAlives)
			},
		},
		{
			name: "no-keepalive with keepalive-time disables keep-alives",
			flags: map[string]interface{}{
				"no-keepalive":   true,
				"keepalive-time": 30,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.True(t, transport.DisableKeepAlives)
			},
		},
		{
			name: "connect-timeout and keepalive-time both configure the dialer",
			flags: map[string]interface{}{
				"connect-timeout": 5.0,
				"no-keepalive":    false,
				"keepalive-time":  30,
			},
			assertFn: func(t *testing.T, client *http.Client) {
				transport, ok := client.Transport.(*http.Transport)
				require.True(t, ok)
				assert.NotNil(t, transport.DialContext)
				assert.False(t, transport.DisableKeepAlives)
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

func TestBuildClientTLS(t *testing.T) {
	certPEM, keyPEM := generateTestCert(t)
	certFile := writeTempFile(t, certPEM)
	keyFile := writeTempFile(t, keyPEM)

	t.Run("cacert loads custom CA pool", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("cacert", certFile, "")
		cmd.Flags().String("cert", "", "")
		cmd.Flags().String("key", "", "")
		cmd.Flags().Bool("insecure", false, "")

		client, err := BuildClient(cmd)
		require.NoError(t, err)
		require.NotNil(t, client)

		transport, ok := client.Transport.(*http.Transport)
		require.True(t, ok)
		require.NotNil(t, transport.TLSClientConfig)
		assert.NotNil(t, transport.TLSClientConfig.RootCAs)
	})

	t.Run("cert and key loads client certificate", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("cacert", "", "")
		cmd.Flags().String("cert", certFile, "")
		cmd.Flags().String("key", keyFile, "")
		cmd.Flags().Bool("insecure", false, "")

		client, err := BuildClient(cmd)
		require.NoError(t, err)
		require.NotNil(t, client)

		transport, ok := client.Transport.(*http.Transport)
		require.True(t, ok)
		require.NotNil(t, transport.TLSClientConfig)
		assert.Len(t, transport.TLSClientConfig.Certificates, 1)
	})

	t.Run("cacert with invalid file returns error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("cacert", "/nonexistent/ca.pem", "")
		cmd.Flags().String("cert", "", "")
		cmd.Flags().String("key", "", "")
		cmd.Flags().Bool("insecure", false, "")

		_, err := BuildClient(cmd)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "reading cacert")
	})

	t.Run("cert with invalid file returns error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("cacert", "", "")
		cmd.Flags().String("cert", "/nonexistent/cert.pem", "")
		cmd.Flags().String("key", "/nonexistent/key.pem", "")
		cmd.Flags().Bool("insecure", false, "")

		_, err := BuildClient(cmd)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "loading client certificate")
	})
}
