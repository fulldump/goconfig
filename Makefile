PROJECT = github.com/fulldump/goconfig

GOPATH_FULL=$(shell pwd)/../../../../
GOPATH=$(shell readlink -f $(GOPATH_FULL))
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

.PHONY: all setup test coverage example

all:	test

setup:
	$(GOCMD) get gopkg.in/mgo.v2/bson

test: setup
	$(GOCMD) test $(PROJECT)/... -cover

example: setup
	$(GOCMD) build -o $(GOPATH)/bin/example $(PROJECT)/example

coverage: setup
	rm -fr coverage
	mkdir -p coverage
	$(GOCMD) list $(PROJECT)/... > coverage/packages
	@i=a ; \
	while read -r P; do \
		i=a$$i ; \
		$(GOCMD) test $$P -cover -covermode=count -coverprofile=coverage/$$i.out; \
	done <coverage/packages

	echo "mode: count" > coverage/coverage
	cat coverage/*.out | grep -v "mode: count" >> coverage/coverage
	$(GOCMD) tool cover -html=coverage/coverage
