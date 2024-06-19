.PHONY: test clean build

test:
	go test -v -count=1 ./...

# Clean up any build artifacts
clean:
	go clean

format :
	go fmt ./...