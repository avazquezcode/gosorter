test:
	go test ./...

tools:
	go install github.com/vektra/mockery/v2@v2.46.3

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./internal/...`

gen-docs:
	go run cmd/docs/doc_gen.go
