# Contributing

Thanks for contributing to `p2`.

## Development

1. Fork the repository and create a branch for your change.
2. Keep changes focused and document any behavior changes in `README.md`.
3. Run the test suite before opening a pull request:

```sh
go test ./...
go build ./cmd/p2
```

## Style

- Follow standard Go formatting with `gofmt`.
- Add or update tests when behavior changes.
- Keep the CLI behavior simple and explicit.

## Pull Requests

- Explain the user-visible behavior change clearly.
- Include command output or tests when relevant.
- Keep pull requests small enough to review quickly.
