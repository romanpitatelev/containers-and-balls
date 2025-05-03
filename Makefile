run: 
	go build -o bin/main ./cmd/containers-and-balls/main.go
	./bin/main

tidy:
	go mod tidy

lint: tidy
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	golangci-lint run ./...

test: 
	go test -v ./...