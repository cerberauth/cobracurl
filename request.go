package cobracurl

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

func BuildRequest(cmd *cobra.Command, args []string) (*http.Request, error) {
	method, _ := cmd.Flags().GetString("request")
	rawURL, _ := cmd.Flags().GetString("url")
	if rawURL == "" && len(args) > 0 {
		rawURL = args[0]
	}

	forceGet, _ := cmd.Flags().GetBool("get")
	if forceGet && method == "" {
		method = "GET"
	}
	if method == "" {
		if head, _ := cmd.Flags().GetBool("head"); head {
			method = "HEAD"
		}
	}

	if method == "" || rawURL == "" {
		return nil, ErrMissingRequiredFields
	}

	var body string
	var extraHeaders []string

	data, _ := cmd.Flags().GetString("data")
	dataBinary, _ := cmd.Flags().GetString("data-binary")
	dataRaw, _ := cmd.Flags().GetString("data-raw")
	dataUrlencode, _ := cmd.Flags().GetString("data-urlencode")

	switch {
	case data != "":
		body = data
	case dataBinary != "":
		body = dataBinary
	case dataRaw != "":
		body = dataRaw
	case dataUrlencode != "":
		body = encodeData(dataUrlencode)
	default:
		if formMap, _ := cmd.Flags().GetStringToString("form"); len(formMap) > 0 {
			formData := make([]string, 0, len(formMap))
			for k, v := range formMap {
				formData = append(formData, k+"="+v)
			}
			body = strings.Join(formData, "&")
			extraHeaders = append(extraHeaders, "Content-Type: application/x-www-form-urlencoded")
		} else if jsonData, _ := cmd.Flags().GetString("json"); jsonData != "" {
			body = jsonData
			extraHeaders = append(extraHeaders, "Content-Type: application/json")
			extraHeaders = append(extraHeaders, "Accept: application/json")
		}
	}

	if forceGet && body != "" {
		separator := "?"
		if strings.Contains(rawURL, "?") {
			separator = "&"
		}
		rawURL = rawURL + separator + body
		body = ""
		extraHeaders = nil
	}

	var requestBody *bytes.Reader
	if body != "" {
		requestBody = bytes.NewReader([]byte(body))
	} else {
		requestBody = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(strings.ToUpper(method), rawURL, requestBody)
	if err != nil {
		return nil, err
	}

	if compressed, _ := cmd.Flags().GetBool("compressed"); compressed {
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	}

	if rangeVal, _ := cmd.Flags().GetString("range"); rangeVal != "" {
		req.Header.Set("Range", "bytes="+rangeVal)
	}

	if userAgent, _ := cmd.Flags().GetString("user-agent"); userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}

	if userArg, _ := cmd.Flags().GetString("user"); userArg != "" {
		parts := strings.SplitN(userArg, ":", 2)
		if len(parts) == 2 {
			req.SetBasicAuth(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}

	if bearer, _ := cmd.Flags().GetString("oauth2-bearer"); bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}

	if referer, _ := cmd.Flags().GetString("referer"); referer != "" {
		req.Header.Set("Referer", referer)
	}

	headers, _ := cmd.Flags().GetStringArray("header")
	for _, h := range append(extraHeaders, headers...) {
		if h == "" {
			continue
		}
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 {
			req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}

	cookies, _ := cmd.Flags().GetStringArray("cookie")
	for _, cookieStr := range cookies {
		for _, pair := range strings.Split(cookieStr, ";") {
			pair = strings.TrimSpace(pair)
			if pair == "" {
				continue
			}
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				req.AddCookie(&http.Cookie{
					Name:  strings.TrimSpace(parts[0]),
					Value: strings.TrimSpace(parts[1]),
				})
			}
		}
	}

	return req, nil
}

func encodeData(s string) string {
	if idx := strings.IndexByte(s, '='); idx >= 0 {
		return s[:idx+1] + url.QueryEscape(s[idx+1:])
	}
	return url.QueryEscape(s)
}

var ErrMissingRequiredFields = errors.New("missing required fields: method and url")
