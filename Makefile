test:
	go test ./...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./internal/...`
