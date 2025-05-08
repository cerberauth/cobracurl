# cobracurl

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cerberauth/cobracurl/ci.yml?branch=main&label=core%20build&style=for-the-badge)](https://github.com/cerberauth/cobracurl/actions/workflows/ci.yml)
![Latest version](https://img.shields.io/github/v/release/cerberauth/cobracurl?sort=semver&style=for-the-badge)
[![Github Repo Stars](https://img.shields.io/github/stars/cerberauth/cobracurl?style=for-the-badge)](https://github.com/cerberauth/cobracurl)
![License](https://img.shields.io/github/license/cerberauth/cobracurl?style=for-the-badge)

**cobracurl** is a Go package that helps you build HTTP requests from Cobra CLI commands — perfect for developers converting `curl` commands into reproducible CLI tools.

If you're building a CLI app with [Cobra](https://github.com/spf13/cobra) and want to recreate or debug HTTP requests using cURL-like arguments, `cobracurl` provides an easy way to translate CLI flags into a fully-formed `*http.Request`.

## ✨ Features

- Define CLI flags for common HTTP request elements (method, URL, headers, body, etc.)
- Generate `*http.Request` objects from those flags
- Minimal, dependency-free, and composable
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

2. Build the HTTP request in your command's Run function

```go
cmd := &cobra.Command{
    Use: "send",
    RunE: func(cmd *cobra.Command, args []string) error {
        req, err := cobracurl.BuildRequest(cmd)
        if err != nil {
            return err
        }

        client := &http.Client{}
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
  --method POST \
  --url https://api.example.com/data \
  --header "Content-Type: application/json" \
  --data '{"foo":"bar"}'
```

## 📦 API

```go
func RegisterFlags(flags *pflag.FlagSet)
```

Register the flags for HTTP method, URL, headers, and body. This function should be called in the `init()` function of your Cobra command.

```go
func BuildRequest(cmd *cobra.Command) (*http.Request, error)
```

Builds an `*http.Request` object based on the flags set in the Cobra command. It returns an error if any required flags are missing or if the request cannot be created.

## Example

See example/ for a minimal CLI tool using cobracurl.

## License

This repository is licensed under the [MIT License](https://github.com/cerberauth/cobracurl/blob/main/LICENSE) @ [CerberAuth](https://www.cerberauth.com/). You are free to use, modify, and distribute the contents of this repository for educational and testing purposes.
