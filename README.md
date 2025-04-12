# Google Play Console Go Client Integration

This repository is a [Go](https://go.dev) implementation for making resource requests to the [Google Play Console](https://developers.google.com/android-publisher/api-ref/rest).

This project is still a work-in-progress, but will be used to enable managing resources using a custom Terraform Provider.

## Usage

See `main.go` for example usage.

The main method can be called using:

```bash
go run main.go "~/path/to/service-account.json" "DEVELOPER_ACCOUNT_ID"
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Requirements

- [Go](https://golang.org/doc/install) >= 1.22

### Building the client

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

### Commits

[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) are required for each pull request to ensure that release versioning can be managed automatically.
Please ensure that you have enabled the Git hooks, so that you don't get caught out!:
```
git config core.hooksPath hooks
```