build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go 
deploy: build
	serverless deploy --aws-profile serverlessEhsan