.PHONY: build
build:
	go build -o buffnet ./cmd/cli

.PHONY: test
test:
	go test ./...