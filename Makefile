test:
	go test -timeout 30s -race -cover ./...

wails:
	go build -tags wails ./...

cov:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out