build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go 
test: 
	go test -v ./pkg/utils