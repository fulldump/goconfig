GOCMD=go

.PHONY: all test coverage example

all: test

test:
	$(GOCMD) version
	$(GOCMD) test -cover ./...

example:
	$(GOCMD) run ./example -help

coverage:
	$(GOCMD) test ./... -cover -covermode=count -coverprofile=coverage.out; \
	$(GOCMD) tool cover -html=coverage.out
