run:
	go run main.go sort words.txt --unique --top=5 --algorithm=mergesort

test:
	go test ./...

tools:
	go install github.com/vektra/mockery/v2@v2.46.3