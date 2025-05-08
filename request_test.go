package cobracurl

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestBuildRequest(t *testing.T) {
	tests := []struct {
		name            string
		flags           map[string]interface{}
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
				"method": "GET",
				"url":    "http://example.com",
			},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
		},
		{
			name:          "Missing method and URL",
			flags:         map[string]interface{}{},
			expectedError: ErrMissingRequiredFields,
		},
		{
			name: "POST request with body and headers",
			flags: map[string]interface{}{
				"method": "POST",
				"url":    "http://example.com",
				"body":   "key=value",
				"header": []string{"Content-Type: application/json", "Authorization: Bearer token"},
			},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "key=value",
			expectedHeaders: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "Bearer token",
			},
		},
		{
			name: "Form data with headers",
			flags: map[string]interface{}{
				"method": "POST",
				"url":    "http://example.com",
				"form":   map[string]string{"key1": "value1", "key2": "value2"},
			},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "POST",
			expectedBody:   "key1=value1&key2=value2",
			expectedHeaders: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		},
		{
			name: "Request with cookies",
			flags: map[string]interface{}{
				"method": "GET",
				"url":    "http://example.com",
				"cookie": []string{"session=abc123", "user=admin"},
			},
			expectedError:  nil,
			expectedURL:    "http://example.com",
			expectedMethod: "GET",
			expectedCookies: map[string]string{
				"session": "abc123",
				"user":    "admin",
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
				case []string:
					cmd.Flags().StringArray(key, v, "")
				case map[string]string:
					cmd.Flags().StringToString(key, v, "")
				}
			}

			req, err := BuildRequest(cmd)

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
					if body.String() != tt.expectedBody {
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
