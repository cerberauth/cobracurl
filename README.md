# cobracurl

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cerberauth/cobracurl/ci.yml?branch=main&label=core%20build&style=for-the-badge)](https://github.com/cerberauth/cobracurl/actions/workflows/ci.yml)
![Latest version](https://img.shields.io/github/v/release/cerberauth/cobracurl?sort=semver&style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/gh/cerberauth/cobracurl?token=BD1WPXJDAW&style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/cerberauth/cobracurl?style=for-the-badge)](https://goreportcard.com/report/github.com/cerberauth/cobracurl)
[![GoDoc reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/cerberauth/cobracurl)

**cobracurl** is a Go package that helps you build HTTP requests from Cobra CLI commands — perfect for developers converting `curl` commands into reproducible CLI tools.

If you're building a CLI app with [Cobra](https://github.com/spf13/cobra) and want to recreate or debug HTTP requests using cURL-like arguments, `cobracurl` provides an easy way to translate CLI flags into a fully-formed `*http.Request`.

## Features

- Define CLI flags for common HTTP request elements (method, URL, headers, body, etc.)
- Generate `*http.Request` objects from those flags
- Generate a pre-configured `*http.Client` (TLS, redirects, timeouts, proxy)
- Rate-limit requests via `--rate` using a standard `*rate.Limiter`
- Minimal and composable
- Easy integration with existing Cobra commands

## 🔧 Installation

```bash
go get github.com/cerberauth/cobracurl
```

## Usage

1. Register flags on your Cobra command

```go
import "github.com/cerberauth/cobracurl"

func init() {
    cobracurl.RegisterFlags(rootCmd.Flags())
}
```

2. Build the HTTP request, client, and optional rate limiter in your command's Run function

```go
cmd := &cobra.Command{
    Use: "send",
    RunE: func(cmd *cobra.Command, args []string) error {
        req, err := cobracurl.BuildRequest(cmd, args)
        if err != nil {
            return err
        }

        client, err := cobracurl.BuildClient(cmd)
        if err != nil {
            return err
        }

        rl, err := cobracurl.BuildRateLimiter(cmd)
        if err != nil {
            return err
        }

        if rl != nil {
            if err := rl.Wait(cmd.Context()); err != nil {
                return err
            }
        }

        resp, err := client.Do(req)
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        body, _ := io.ReadAll(resp.Body)
        fmt.Println("Response:", string(body))
        return nil
    },
}
```

3. Example CLI command

```bash
yourcli send \
  --request POST \
  --url https://api.example.com/data \
  --header "Content-Type: application/json" \
  --data '{"foo":"bar"}' \
  --location \
  --insecure \
  --rate 10/s
```

## API

```go
func RegisterFlags(flags *pflag.FlagSet)
```

Registers all supported curl-compatible flags on the given flag set. Call this in the `init()` function of your Cobra command.

```go
func BuildRequest(cmd *cobra.Command, args []string) (*http.Request, error)
```

Builds an `*http.Request` from the flags set on the command. The first positional argument is used as the URL if `--url` is not set; a missing scheme defaults to `http://` (see `NormalizeURL`). Returns an error if `--request` and URL are both missing.

Supported flags include: `--request`/`-X`, `--url`, `--header`/`-H`, `--data`/`-d`, `--data-binary`, `--data-raw`, `--data-urlencode`, `--form`/`-F`, `--json`, `--user`/`-u`, `--oauth2-bearer`, `--user-agent`/`-A`, `--referer`/`-e`, `--cookie`/`-b`, `--head`/`-I`, `--get`/`-G`, `--compressed`, `--range`/`-r`.

```go
func NormalizeURL(raw string) (string, error)
```

Adds a default scheme to a URL that has none: `http://`, or `https://` when the port is 443. It handles `host`, `host:port`, `host:port/path` and IPv6 literals, and returns URLs that already have a `scheme://` unchanged. For example `localhost:5000` becomes `http://localhost:5000` (where `url.Parse` would read `localhost` as the scheme) and `[::1]:8080/x` becomes `http://[::1]:8080/x`. It does no network probe. `BuildRequest` applies it to the URL, and it is exported so other commands can normalize a bare host.

```go
func BuildClient(cmd *cobra.Command, opts ...ClientOption) (*http.Client, error)
```

Builds an `*http.Client` from the flags set on the command. Unlike the default Go HTTP client, redirects are **not** followed unless `--location` is set, matching curl's default behavior. The transport is cloned from `http.DefaultTransport`, so `HTTP_PROXY`/`HTTPS_PROXY`/`NO_PROXY` are honored when `--proxy` is unset.

Like curl, no overall timeout is applied by default. Embedders can set a fallback with `cobracurl.WithDefaultTimeout(d)`; `--max-time` still takes precedence.

Supported flags include: `--insecure`/`-k`, `--location`/`-L`, `--max-redirs`, `--max-time`/`-m`, `--connect-timeout`, `--proxy`/`-x`, `--noproxy`.

`--noproxy` takes a comma-separated list of hosts that bypass the proxy, as in curl: a host name matches itself and its subdomains, IP addresses and CIDR ranges are accepted, and `*` bypasses the proxy for every host.

```go
func BuildRateLimiter(cmd *cobra.Command) (*rate.Limiter, error)
```

Parses the `--rate` flag and returns a `*rate.Limiter` (from `golang.org/x/time/rate`) configured with burst=1 for a steady, non-bursty rate. Returns `nil` if `--rate` is not set, meaning no rate limiting is applied. Call `rl.Wait(ctx)` before each request. A rate without a unit is applied per hour, as in curl, and a warning is written to the command's stderr because it is easy to misread as per second.

The flag accepts curl-style rate strings:

| Value | Meaning |
|-------|---------|
| `10/s` | 10 requests per second |
| `100/m` | 100 requests per minute |
| `1000/h` | 1 000 requests per hour |
| `5000/d` | 5 000 requests per day |
| `60` | 60 requests per hour (default unit when omitted, with a warning) |

```go
func ParseRate(rateStr string) (limiter *rate.Limiter, warning string, err error)
```

Lower-level helper that parses a rate string directly without a cobra command. `warning` is non-empty when the rate has no unit (applied per hour); the limiter is still returned. `ParseRate("10/s")` returns an empty warning.

## Example

See [example/](/example/README.md) for a minimal CLI tool using cobracurl.

## License

This repository is licensed under the [MIT License](https://github.com/cerberauth/cobracurl/blob/main/LICENSE) @ [CerberAuth](https://www.cerberauth.com/).
