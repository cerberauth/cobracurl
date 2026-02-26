//go:build integration

package cobracurl

// Integration tests that compare BuildRequest output against real curl.
//
// Run with: go test -tags integration -run TestCompatWithCurl ./...
//
// Each test case fires the same logical HTTP request through two paths:
//   1. Real curl (exec.Command) → echo server → capturedRequest A
//   2. BuildRequest + http.Client  → echo server → capturedRequest B
// Then A and B are compared for method, body, headers and cookies.
//
// Note: flags are registered manually with the types BuildRequest actually
// reads via GetString/GetStringArray/GetStringToString.  RegisterFlags has
// type mismatches for "header", "cookie", "form" and the body flag ("body"
// vs "data") that prevent those features from working end-to-end through the
// public API today.

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// alwaysSkipHeaders lists headers whose default values differ between curl and
// cobracurl, so they are excluded from the general header comparison.
var alwaysSkipHeaders = map[string]bool{
	"User-Agent":      true, // curl sends "curl/x.x" by default; cobracurl sends nothing
	"Accept":          true, // curl sends "*/*" by default
	"Accept-Encoding": true, // Go's http.Client adds "gzip"; curl does not
	"Content-Length":  true, // auto-calculated; may differ for empty GET bodies
}

type capturedRequest struct {
	Method  string
	Headers http.Header
	Body    string
	Cookies map[string]string
}

func newEchoServer(t *testing.T) (*httptest.Server, func() capturedRequest) {
	t.Helper()
	ch := make(chan capturedRequest, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		cookies := make(map[string]string)
		for _, c := range r.Cookies() {
			cookies[c.Name] = c.Value
		}
		ch <- capturedRequest{
			Method:  r.Method,
			Headers: r.Header.Clone(),
			Body:    string(body),
			Cookies: cookies,
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)
	return srv, func() capturedRequest {
		t.Helper()
		select {
		case req := <-ch:
			return req
		case <-time.After(5 * time.Second):
			t.Fatal("timeout: no request received by echo server")
			return capturedRequest{}
		}
	}
}

func runCurl(t *testing.T, url string, args ...string) {
	t.Helper()
	cmdArgs := append([]string{"--silent", "--output", "/dev/null"}, args...)
	cmdArgs = append(cmdArgs, url)
	out, err := exec.Command("curl", cmdArgs...).CombinedOutput()
	require.NoError(t, err, "curl failed: %s", out)
}

// buildAndRun creates a Cobra command with the given flags, builds the request
// via BuildRequest, and fires it against url.
//
// Supported flag value types:
//   - string       → cmd.Flags().Set(name, value)
//   - []string     → cmd.Flags().Set(name, value) called once per element (StringArray)
//   - map[string]string → packed as "k=v,k2=v2" for StringToString flags
func buildAndRun(t *testing.T, url string, flags map[string]interface{}) {
	t.Helper()
	cmd := &cobra.Command{}
	// Register with the types BuildRequest actually reads.
	cmd.Flags().String("request", "", "")
	cmd.Flags().String("url", "", "")
	cmd.Flags().String("body", "", "")
	cmd.Flags().StringArray("header", nil, "")
	cmd.Flags().StringArray("cookie", nil, "")
	cmd.Flags().StringToString("form", nil, "")
	cmd.Flags().String("user-agent", "", "")
	cmd.Flags().String("user", "", "")

	require.NoError(t, cmd.Flags().Set("url", url))
	for name, val := range flags {
		switch v := val.(type) {
		case string:
			require.NoError(t, cmd.Flags().Set(name, v))
		case []string:
			for _, s := range v {
				require.NoError(t, cmd.Flags().Set(name, s))
			}
		case map[string]string:
			pairs := make([]string, 0, len(v))
			for k, mv := range v {
				pairs = append(pairs, k+"="+mv)
			}
			require.NoError(t, cmd.Flags().Set(name, strings.Join(pairs, ",")))
		}
	}

	req, err := BuildRequest(cmd, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	resp.Body.Close()
}

func assertRequestsMatch(t *testing.T, curlReq, cobraReq capturedRequest, extraSkip map[string]bool) {
	t.Helper()
	skip := func(h string) bool {
		return alwaysSkipHeaders[h] || extraSkip[h]
	}

	assert.Equal(t, curlReq.Method, cobraReq.Method, "HTTP method")
	assertBodiesEqual(t, curlReq.Body, cobraReq.Body)
	assert.Equal(t, curlReq.Cookies, cobraReq.Cookies, "cookies")

	// Every non-skipped header sent by curl must also appear in cobracurl.
	for k, curlVals := range curlReq.Headers {
		if skip(k) {
			continue
		}
		cobraVals, ok := cobraReq.Headers[k]
		if assert.True(t, ok, "header %q present in curl but absent in cobracurl", k) {
			assert.Equal(t, curlVals, cobraVals, "header %q value", k)
		}
	}

	// Every non-skipped header sent by cobracurl must also appear in curl.
	for k := range cobraReq.Headers {
		if skip(k) {
			continue
		}
		if _, ok := curlReq.Headers[k]; !ok {
			t.Errorf("header %q present in cobracurl but absent in curl", k)
		}
	}
}

func TestCompatWithCurl(t *testing.T) {
	if _, err := exec.LookPath("curl"); err != nil {
		t.Skip("curl not found in PATH")
	}

	tests := []struct {
		name       string
		curlArgs   []string
		cobraFlags map[string]interface{}
		extraSkip  map[string]bool
		// verify runs after assertRequestsMatch for additional assertions
		// (e.g. checking header values that are in alwaysSkipHeaders).
		verify func(*testing.T, capturedRequest, capturedRequest)
	}{
		{
			name:     "GET request",
			curlArgs: []string{"-X", "GET"},
			cobraFlags: map[string]interface{}{
				"request": "GET",
			},
		},
		{
			name: "POST with JSON body",
			curlArgs: []string{
				"-X", "POST",
				"-H", "Content-Type: application/json",
				"-d", `{"key":"value"}`,
			},
			cobraFlags: map[string]interface{}{
				"request": "POST",
				"header":  []string{"Content-Type: application/json"},
				"body":    `{"key":"value"}`,
			},
		},
		{
			name:     "POST with basic auth",
			curlArgs: []string{"-X", "POST", "-u", "alice:s3cr3t"},
			cobraFlags: map[string]interface{}{
				"request": "POST",
				"user":    "alice:s3cr3t",
			},
		},
		{
			name:     "GET with custom header",
			curlArgs: []string{"-X", "GET", "-H", "X-Request-ID: abc123"},
			cobraFlags: map[string]interface{}{
				"request": "GET",
				"header":  []string{"X-Request-ID: abc123"},
			},
		},
		{
			name:     "GET with cookie",
			curlArgs: []string{"-X", "GET", "-b", "session=abc123"},
			cobraFlags: map[string]interface{}{
				"request": "GET",
				"cookie":  []string{"session=abc123"},
			},
		},
		{
			name: "POST with multiple headers",
			curlArgs: []string{
				"-X", "POST",
				"-H", "Content-Type: application/json",
				"-H", "X-Trace-ID: xyz",
			},
			cobraFlags: map[string]interface{}{
				"request": "POST",
				"header":  []string{"Content-Type: application/json", "X-Trace-ID: xyz"},
			},
		},
		{
			// User-Agent is in alwaysSkipHeaders so assertRequestsMatch won't
			// compare it; we verify the value explicitly in verify() instead.
			name:     "GET with custom user agent",
			curlArgs: []string{"-X", "GET", "-A", "mybot/1.0"},
			cobraFlags: map[string]interface{}{
				"request":    "GET",
				"user-agent": "mybot/1.0",
			},
			verify: func(t *testing.T, curlReq, cobraReq capturedRequest) {
				assert.Equal(t, "mybot/1.0", curlReq.Headers.Get("User-Agent"), "curl User-Agent")
				assert.Equal(t, "mybot/1.0", cobraReq.Headers.Get("User-Agent"), "cobracurl User-Agent")
			},
		},
		{
			// curl with -d implies POST and adds Content-Type: application/x-www-form-urlencoded.
			// cobracurl's form handling produces the same Content-Type and body.
			name:     "POST with form field",
			curlArgs: []string{"-d", "field=hello"},
			cobraFlags: map[string]interface{}{
				"request": "POST",
				"form":    map[string]string{"field": "hello"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, receive := newEchoServer(t)

			runCurl(t, srv.URL+"/test", tt.curlArgs...)
			curlReq := receive()

			buildAndRun(t, srv.URL+"/test", tt.cobraFlags)
			cobraReq := receive()

			assertRequestsMatch(t, curlReq, cobraReq, tt.extraSkip)
			if tt.verify != nil {
				tt.verify(t, curlReq, cobraReq)
			}
		})
	}
}

// assertBodiesEqual compares two request bodies. For URL-encoded bodies it
// parses them into url.Values so that key order does not affect the result.
func assertBodiesEqual(t *testing.T, curlBody, cobraBody string) {
	t.Helper()
	va, errA := url.ParseQuery(curlBody)
	vb, errB := url.ParseQuery(cobraBody)
	if errA == nil && errB == nil && len(va) > 0 && len(vb) > 0 {
		assert.Equal(t, va, vb, "request body (parsed as form values)")
		return
	}
	assert.Equal(t, curlBody, cobraBody, "request body")
}
