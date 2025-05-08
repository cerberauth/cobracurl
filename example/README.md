# cobracurl-demo

**cobracurl-demo** is a Go project that demonstrates the usage of the [`cobracurl`](https://github.com/cerberauth/cobracurl) package to build and send HTTP requests from a Cobra CLI application. This project serves as a practical example for developers looking to integrate HTTP request capabilities into their CLI tools.

## ✨ Features

- Build HTTP requests using CLI flags
- Supports various HTTP methods and headers
- Easy integration with the `cobracurl` package

## 🔧 Installation

To run this project, ensure you have Go installed and then clone the repository:

```bash
git clone https://github.com/cerberauth/cobracurl
cd example
go mod tidy
```

## 🚀 Usage

1. Build the project:

```bash
go build -o yourcli ./cmd
```

2. Run the CLI command to send an HTTP request:

```bash
./yourcli send \
  --method POST \
  --url https://api.example.com/data \
  --header "Content-Type: application/json" \
  --data '{"foo":"bar"}'
```

## 📦 API

This project utilizes the following functions from the `cobracurl` package:

```go
func RegisterFlags(flags *pflag.FlagSet)
```

Register the flags for HTTP method, URL, headers, and body.

```go
func BuildRequest(cmd *cobra.Command) (*http.Request, error)
```

Builds an `*http.Request` object based on the flags set in the Cobra command.

## License

This project is licensed under the [MIT License](https://github.com/cerberauth/cobracurl/blob/main/LICENSE). You are free to use, modify, and distribute the contents of this repository for educational and testing purposes.
