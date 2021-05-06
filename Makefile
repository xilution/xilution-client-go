generate:
	go generate ./...

test: 
	go test ./...

test_coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

format:
	go fmt ./...
