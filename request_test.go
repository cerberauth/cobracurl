package cobracurl

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestBuildRequest(t *testing.T) {
	tests := []struct {
		name            string
		flags           map[string]interface{}
		args            []string
		expectedError   error
		expectedURL     string
		expectedMethod  string
		expectedBody    string
		expectedHeaders map[string]string
		expectedCookies map[string]string
	}{
		{
			name: "Valid GET request",
			flags: map[string]interface{}{
				"request": "GET",
				"url":     "http://example.com",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
		},
		{
			name: "Valid GET request with URL from args",
			flags: map[string]interface{}{
				"request": "GET",
			},
			args:           []string{"http://example.com"},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
		},
		{
			name: "Valid POST request",
			flags: map[string]interface{}{
				"request": "POST",
				"url":     "http://example.com",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
		},
		{
			name:          "Missing method and URL",
			flags:         map[string]interface{}{},
			args:          []string{},
			expectedError: ErrMissingRequiredFields,
		},
		{
			name: "HEAD request via --head flag",
			flags: map[string]interface{}{
				"head": true,
				"url":  "http://example.com",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "HEAD",
		},
		{
			name: "--request overrides --head",
			flags: map[string]interface{}{
				"request": "GET",
				"head":    true,
				"url":     "http://example.com",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
		},
		{
			name: "POST request with data and header",
			flags: map[string]interface{}{
				"request": "POST",
				"url":     "http://example.com",
				"data":    "key=value",
				"header":  []string{"Content-Type: application/json"},
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "key=value",
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "POST request with data-binary",
			flags: map[string]interface{}{
				"request":     "POST",
				"url":         "http://example.com",
				"data-binary": "\x00\x01binary",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "\x00\x01binary",
		},
		{
			name: "POST request with data-raw",
			flags: map[string]interface{}{
				"request":  "POST",
				"url":      "http://example.com",
				"data-raw": "@value",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "@value",
		},
		{
			name: "POST request with data-urlencode (value only)",
			flags: map[string]interface{}{
				"request":        "POST",
				"url":            "http://example.com",
				"data-urlencode": "hello world",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "hello+world",
		},
		{
			name: "POST request with data-urlencode (name=value)",
			flags: map[string]interface{}{
				"request":        "POST",
				"url":            "http://example.com",
				"data-urlencode": "q=hello world",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "q=hello+world",
		},
		{
			name: "GET with data via --get appends to URL",
			flags: map[string]interface{}{
				"get":  true,
				"url":  "http://example.com",
				"data": "key=value",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com?key=value",
			expectedMethod: "GET",
			expectedBody:   "",
		},
		{
			name: "GET with data via --get appends to existing query string",
			flags: map[string]interface{}{
				"get":  true,
				"url":  "http://example.com?existing=1",
				"data": "key=value",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com?existing=1&key=value",
			expectedMethod: "GET",
		},
		{
			name: "POST request with form data",
			flags: map[string]interface{}{
				"request": "POST",
				"url":     "http://example.com",
				"form":    map[string]string{"key": "value"},
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "key=value",
			expectedHeaders: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		},
		{
			name: "POST request with JSON",
			flags: map[string]interface{}{
				"request": "POST",
				"url":     "http://example.com",
				"json":    `{"key":"value"}`,
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   `{"key":"value"}`,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
		},
		{
			name: "--data takes priority over --json",
			flags: map[string]interface{}{
				"request": "POST",
				"url":     "http://example.com",
				"data":    "raw=body",
				"json":    `{"key":"value"}`,
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "raw=body",
		},
		{
			name: "Compressed request adds Accept-Encoding",
			flags: map[string]interface{}{
				"request":    "GET",
				"url":        "http://example.com",
				"compressed": true,
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"Accept-Encoding": "gzip, deflate, br",
			},
		},
		{
			name: "Range request adds Range header",
			flags: map[string]interface{}{
				"request": "GET",
				"url":     "http://example.com",
				"range":   "0-499",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"Range": "bytes=0-499",
			},
		},
		{
			name: "Request with cookies",
			flags: map[string]interface{}{
				"request": "GET",
				"url":     "http://example.com",
				"cookie":  "session=abc123; user=admin",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedCookies: map[string]string{
				"session": "abc123",
				"user":    "admin",
			},
		},
		{
			name: "Request with user agent",
			flags: map[string]interface{}{
				"request":    "GET",
				"url":        "http://example.com",
				"user-agent": "MyUserAgent/1.0",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"User-Agent": "MyUserAgent/1.0",
			},
		},
		{
			name: "Request with basic auth",
			flags: map[string]interface{}{
				"request": "GET",
				"url":     "http://example.com",
				"user":    "username:password",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"Authorization": "Basic dXNlcm5hbWU6cGFzc3dvcmQ=",
			},
		},
		{
			name: "Request with OAuth2 bearer token",
			flags: map[string]interface{}{
				"request":       "GET",
				"url":           "http://example.com",
				"oauth2-bearer": "mytoken",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"Authorization": "Bearer mytoken",
			},
		},
		{
			name: "Request with referer",
			flags: map[string]interface{}{
				"request": "GET",
				"url":     "http://example.com",
				"referer": "http://referrer.example.com",
			},
			args:           []string{},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedHeaders: map[string]string{
				"Referer": "http://referrer.example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for key, value := range tt.flags {
				switch v := value.(type) {
				case string:
					cmd.Flags().String(key, v, "")
				case bool:
					cmd.Flags().Bool(key, v, "")
				case []string:
					cmd.Flags().StringArray(key, v, "")
				case map[string]string:
					cmd.Flags().StringToString(key, v, "")
				}
			}

			req, err := BuildRequest(cmd, tt.args)

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}

			if err == nil {
				if req.URL.String() != tt.expectedURL {
					t.Errorf("expected URL %s, got %s", tt.expectedURL, req.URL.String())
				}

				if req.Method != tt.expectedMethod {
					t.Errorf("expected method %s, got %s", tt.expectedMethod, req.Method)
				}

				if tt.expectedBody != "" {
					body := new(strings.Builder)
					_, _ = io.Copy(body, req.Body)
					if !bodiesEqual(body.String(), tt.expectedBody) {
						t.Errorf("expected body %s, got %s", tt.expectedBody, body.String())
					}
				}

				for key, value := range tt.expectedHeaders {
					if req.Header.Get(key) != value {
						t.Errorf("expected header %s: %s, got %s", key, value, req.Header.Get(key))
					}
				}

				for _, cookie := range req.Cookies() {
					if tt.expectedCookies[cookie.Name] != cookie.Value {
						t.Errorf("expected cookie %s: %s, got %s", cookie.Name, tt.expectedCookies[cookie.Name], cookie.Value)
					}
				}
			}
		})
	}
}

func bodiesEqual(a, b string) bool {
	va, errA := url.ParseQuery(a)
	vb, errB := url.ParseQuery(b)
	if errA == nil && errB == nil && len(va) > 0 && len(vb) > 0 {
		if len(va) != len(vb) {
			return false
		}
		for k, vals := range va {
			bVals, ok := vb[k]
			if !ok || len(vals) != len(bVals) {
				return false
			}
			for i := range vals {
				if vals[i] != bVals[i] {
					return false
				}
			}
		}
		return true
	}
	return a == b
}
