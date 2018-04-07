PROJECT        = goconfig
GOLANG_VERSION = 1.9
VERSION        = $(shell git describe --tags --always)

DIR_TOOLS      = $(CURDIR)/src/tools
DIR_TESTREPORT = test

GOPATH       = $(CURDIR)
GO			 = GOPATH=$(GOPATH) go
GOLINT     	 = $(GO) run $(DIR_TOOLS)/vendor/github.com/golang/lint/golint/golint.go $(DIR_TOOLS)/vendor/github.com/golang/lint/golint/import.go
GOVET     	 = $(GO) tool vet
GO2XUNIT   	 = $(GO) run $(DIR_TOOLS)/vendor/github.com/tebeka/go2xunit/cmdline.go $(DIR_TOOLS)/vendor/github.com/tebeka/go2xunit/main.go
GLIDE        = GOPATH=$(GOPATH) glide update; GOPATH=$(GOPATH) glide install --force
COMMAND_ARGS = cover=$(cover) junit=$(junit) forced=$(forced) vendors=$(vendors)

PACKAGES      = $(shell $(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/)
PACKAGES_TEST = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{.ImportPath}}{{ end }}' ./src/$(PROJECT)/... | grep -v /vendor/)

DOCKER_RUN = docker run -v $(GOPATH):/go --workdir /go golang:$(GOLANG_VERSION)

.PHONY: all
all: clean deps fmt lint test ; @ ## All

.PHONY: setup
setup:
	mkdir -p src
	rm -fr src/*
	ln -s .. src/$(PROJECT)
	ln -s ../tools src/tools

.PHONY: clean ; @ ## Delete temporary directories [vendors:0|1]
vendors?=0
clean: ; @ ## Delete temporary directories [vendors:0|1]
	rm -rf $(DIR_TESTREPORT)
	rm -rf bin
	rm -rf pkg
	rm -rf src
ifeq ($(vendors),1)
	rm -rf $(CURDIR)/vendor
	rm -rf $(CURDIR)/tools/vendor
endif

# TODO glide must be run from code
.PHONY: deps
deps: | setup ; @ ## Download required Golang libraries for project & Tools
	cd $(CURDIR)/src/$(PROJECT); $(GLIDE)
	cd $(DIR_TOOLS); $(GLIDE)

.PHONY: fmt
fmt: | setup ; @ ## Code formatter
	for pkg in $(PACKAGES); do \
		gofmt -l -w  -e $$pkg/*.go; \
	done

.PHONY: lint
forced ?= 0
lint: | setup ; @ ## Code analysis [forced:0|1]
ifeq ($(forced),1)
	-for pkg in $(PACKAGES); do \
   		$(GOVET) $$pkg/*.go; \
    done;
	@echo ""
	$(GOLINT) -set_exit_status ./...
else
	-for pkg in $(PACKAGES); do \
    	$(GOVET) $$pkg/*.go; \
    done;
	@echo ""
	-$(GOLINT) $(PACKAGES)
endif

.PHONY: test
cover ?= 0
timeout ?=20
junit ?=0
testArgs := -v -timeout $(timeout)s -short
ifeq ($(cover),1)
testArgs += -cover
endif
ifeq ($(junit),1)
testArgs += | tee $(DIR_TESTREPORT)/$(PROJECT)-test.output
endif
test: | setup ; @ ## Run tests [junit:0|1, cover:0|1, timeout:<seconds>]
ifeq ($(junit),1)
	mkdir -p test
endif
	$(GO) test -p=1 $(PACKAGES_TEST)  ${testArgs}; \
	status=$$?; \
	exit $$status
ifeq ($(junit),1)
	$(GO2XUNIT) -fail -input $(DIR_TESTREPORT)/$(PROJECT)-test.output -output $(DIR_TESTREPORT)/$(PROJECT)-test.xml
endif

.PHONY: docker-%
docker-%: ; @ ## Run commands from docker
	$(DOCKER_RUN) make $* $(COMMAND_ARGS)

.PHONY: version
version: ; @ ## Current version
	@printf "version $(VERSION)\n"

.PHONY: help
help:
	@printf "\n\033[0;31m---------------------------------\n"
	@printf "\033[0;37m    The Golang Makefile\n"
	@printf "\033[0;31m---------------------------------\n\n"
	@printf "\033[0;31mCommands\n\n"
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-15s\033[0m %s\n", $$1,$$2}'
	@printf "\n"