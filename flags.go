package cobracurl

import (
	"github.com/spf13/pflag"
)

func RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolP("append", "a", false, "Append to target file(s) instead of overwriting")
	flags.StringP("cookie", "b", "", "Send cookies from string/file")
	flags.BoolP("basic", "", false, "Use HTTP Basic Auth (use with -u)")
	flags.BoolP("compressed", "", false, "Request compressed response")
	flags.StringP("data", "d", "", "HTTP POST data")
	flags.BoolP("digest", "", false, "Use HTTP Digest Auth (use with -u)")
	flags.BoolP("fail", "f", false, "Fail fast with no output on HTTP errors")
	flags.StringP("form", "F", "", "Specify multipart MIME data")
	flags.StringP("head", "I", "", "Show document info only")
	flags.StringP("header", "H", "", "Pass custom header(s) to server")
	flags.StringP("get", "G", "", "Put the post data in the URL and use GET")
	flags.BoolP("include", "i", false, "Include protocol response headers in the output")
	flags.BoolP("insecure", "k", false, "Allow insecure server connections when using SSL")
	flags.StringP("json", "", "", "HTTP POST JSON")
	flags.StringP("method", "X", "GET", "Specify request method to use")
	flags.StringP("output", "o", "", "Write to file instead of stdout")
	flags.BoolP("location", "L", false, "Follow redirects")
	flags.StringP("proxy", "x", "", "Use this proxy")
	flags.StringP("remote-name", "O", "", "Write output to a file named as the remote file")
	flags.BoolP("silent", "s", false, "Silent mode")
	flags.StringP("referer", "e", "", "Send Referer Page information.")
	flags.StringP("upload-file", "T", "", "Transfer local FILE to destination")
	flags.StringP("url", "", "", "URL to work with")
	flags.StringP("user-agent", "A", "", "Send User-Agent <name> to server")
	flags.StringP("user", "u", "", "Server user and password")
	flags.BoolP("verbose", "v", false, "Make the operation more talkative")
}
