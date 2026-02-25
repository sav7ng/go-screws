# AGENTS.md

## Cursor Cloud specific instructions

This is a **Go utility library** (`github.com/sav7ng/go-screws`), not a runnable service. There is no application to start — development consists of editing library code and running tests.

### Key commands

- **Lint:** `go vet ./...`
- **Test:** `go test ./...`
- **Build:** `go build ./...`
- **Run demo:** `go run examples/hello_world.go`

### Known issues

- `go vet` reports unreachable code in `convert/convert.go:54` — this is a pre-existing repo issue.
- `catch/catch_test.go` always fails because the test panics without using `defer` on `Dmp()`. This is a pre-existing test bug.

### Go version

The `page/common.pb.go` (protobuf-generated) requires `google.golang.org/protobuf`, which needs Go 1.23+. The VM environment installs Go 1.23.6 to satisfy this. If Go is not at 1.23+, run: `wget -q https://go.dev/dl/go1.23.6.linux-amd64.tar.gz -O /tmp/go.tar.gz && sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf /tmp/go.tar.gz && rm /tmp/go.tar.gz`

### No external services needed

This library has zero runtime service dependencies (no databases, no Docker, no network services).
