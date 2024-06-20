test:
	@go test -v -count=1 ./...
	@echo "test completed"

# Clean up any build artifacts
clean:
	@go clean
	@echo "cleaned"

format :
	@go fmt ./...
	@echo "formatted completed"

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
	@echo "proto generated"

.PHONY: proto