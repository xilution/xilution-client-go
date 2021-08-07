generate:
	go generate ./...

test: 
	go test -v ./...

test_coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

format:
	go fmt ./...

clean:
	rm -rf vendor *_mock.go coverage.out
