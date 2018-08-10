generate:
	@go generate ./...
.PHONY: generate

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
