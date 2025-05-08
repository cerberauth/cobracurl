package cobracurl

import (
	"bytes"
	"errors"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func BuildRequest(cmd *cobra.Command) (*http.Request, error) {
	method, _ := cmd.Flags().GetString("method")
	url, _ := cmd.Flags().GetString("url")

	if method == "" || url == "" {
		return nil, ErrMissingRequiredFields
	}

	body, _ := cmd.Flags().GetString("body")
	headers, _ := cmd.Flags().GetStringArray("header")

	form, _ := cmd.Flags().GetStringToString("form")
	if len(form) > 0 {
		formData := make([]string, 0, len(form))
		for key, value := range form {
			formData = append(formData, key+"="+value)
		}
		body = strings.Join(formData, "&")
		headers = append(headers, "Content-Type: application/x-www-form-urlencoded")
	}

	var requestBody *bytes.Reader
	if body != "" {
		requestBody = bytes.NewReader([]byte(body))
	} else {
		requestBody = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, requestBody)
	if err != nil {
		return nil, err
	}

	userAgent, _ := cmd.Flags().GetString("user-agent")
	if userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}

	userArg, _ := cmd.Flags().GetString("user")
	if userArg != "" {
		parts := strings.SplitN(userArg, ":", 2)
		if len(parts) == 2 {
			req.SetBasicAuth(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}

	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) == 2 {
			req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}

	cookies, _ := cmd.Flags().GetStringArray("cookie")
	for _, cookie := range cookies {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 {
			req.AddCookie(&http.Cookie{
				Name:  strings.TrimSpace(parts[0]),
				Value: strings.TrimSpace(parts[1]),
			})
		}
	}

	return req, nil
}

var ErrMissingRequiredFields = errors.New("missing required fields: method and url")
