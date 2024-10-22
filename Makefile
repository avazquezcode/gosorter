test:
	go test ./...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./internal/...`

gen-docs:
	go run cmd/docs/doc_gen.go
