# cobracurl

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cerberauth/cobracurl/ci.yml?branch=main&label=core%20build&style=for-the-badge)](https://github.com/cerberauth/cobracurl/actions/workflows/ci.yml)
![Latest version](https://img.shields.io/github/v/release/cerberauth/cobracurl?sort=semver&style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/gh/cerberauth/cobracurl?token=BD1WPXJDAW&style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/cerberauth/cobracurl?style=for-the-badge)](https://goreportcard.com/report/github.com/cerberauth/cobracurl)
[![GoDoc reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/cerberauth/cobracurl)

**cobracurl** is a Go package that helps you build HTTP requests from Cobra CLI commands — perfect for developers converting `curl` commands into reproducible CLI tools.

If you're building a CLI app with [Cobra](https://github.com/spf13/cobra) and want to recreate or debug HTTP requests using cURL-like arguments, `cobracurl` provides an easy way to translate CLI flags into a fully-formed `*http.Request`.

## ✨ Features

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

## 🚀 Usage

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

## 📦 API

```go
func RegisterFlags(flags *pflag.FlagSet)
```

Registers all supported curl-compatible flags on the given flag set. Call this in the `init()` function of your Cobra command.

```go
func BuildRequest(cmd *cobra.Command, args []string) (*http.Request, error)
```

Builds an `*http.Request` from the flags set on the command. The first positional argument is used as the URL if `--url` is not set. Returns an error if `--request` and URL are both missing.

Supported flags include: `--request`/`-X`, `--url`, `--header`/`-H`, `--data`/`-d`, `--data-binary`, `--data-raw`, `--data-urlencode`, `--form`/`-F`, `--json`, `--user`/`-u`, `--oauth2-bearer`, `--user-agent`/`-A`, `--referer`/`-e`, `--cookie`/`-b`, `--head`/`-I`, `--get`/`-G`, `--compressed`, `--range`/`-r`.

```go
func BuildClient(cmd *cobra.Command) (*http.Client, error)
```

Builds an `*http.Client` from the flags set on the command. Unlike the default Go HTTP client, redirects are **not** followed unless `--location` is set, matching curl's default behavior.

Supported flags include: `--insecure`/`-k`, `--location`/`-L`, `--max-redirs`, `--max-time`/`-m`, `--connect-timeout`, `--proxy`/`-x`.

```go
func BuildRateLimiter(cmd *cobra.Command) (*rate.Limiter, error)
```

Parses the `--rate` flag and returns a `*rate.Limiter` (from `golang.org/x/time/rate`) configured with burst=1 for a steady, non-bursty rate. Returns `nil` if `--rate` is not set, meaning no rate limiting is applied. Call `rl.Wait(ctx)` before each request.

The flag accepts curl-style rate strings:

| Value | Meaning |
|-------|---------|
| `10/s` | 10 requests per second |
| `100/m` | 100 requests per minute |
| `1000/h` | 1 000 requests per hour |
| `5000/d` | 5 000 requests per day |
| `60` | 60 requests per hour (default unit when omitted) |

```go
func ParseRate(rateStr string) (*rate.Limiter, error)
```

Lower-level helper that parses a rate string directly without a cobra command.

## Example

See example/ for a minimal CLI tool using cobracurl.

## License

This repository is licensed under the [MIT License](https://github.com/cerberauth/cobracurl/blob/main/LICENSE) @ [CerberAuth](https://www.cerberauth.com/).
