.PHONY: build release-snapshot run test

build:
	go build ./...

release-snapshot:
	goreleaser release --snapshot --clean

run:
	go run .

test:
	go test ./...
