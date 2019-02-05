PROJECT = github.com/fulldump/goconfig

GOCMD=GOPATH=`pwd` go

.PHONY: all setup test coverage example

all:	test

setup:
	mkdir -p src/$(PROJECT)
	rmdir src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)

test:
	$(GOCMD) version
	$(GOCMD) env
	$(GOCMD) test -v $(PROJECT)

example:
	$(GOCMD) install $(PROJECT)/example

coverage:
	$(GOCMD) test ./src/github.com/fulldump/goconfig -cover -covermode=count -coverprofile=coverage.out; \
	$(GOCMD) tool cover -html=coverage.out

