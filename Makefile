-include .env

VERSION                 := $(shell git describe --tags)
BUILD                   := $(shell git rev-parse --short HEAD)
PROJECTNAME             := $(shell basename "$(PWD)")
IMAGE_NAME              := oojob/company
STDERR                  := /tmp/.$(PROJECTNAME)-stderr.txt # Redirect error output to a file, so we can show it in development mode.
PID                     := /tmp/.$(PROJECTNAME)-api-server.pid # PID file will store the server process id when it's running on development mode
SERVER_OUT              := "bin/server"
ENTRYPOINT              := "entry-point.sh"
PKG                     := "github.com/oojob/company"
SERVER_PKG_BUILD        := "${PKG}"
PROTOC_ZIP              := "protoc-3.11.4-linux-x86_64.zip"
PKG_LIST                := $(shell go list ${PKG}/... | grep -v /vendor/)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

setup-protobuf-linux: ## install protobuff on linux
	@apt-get update && apt-get -y install unzip
	@curl -OL https://github.com/google/protobuf/releases/download/v3.11.4/$(PROTOC_ZIP)
	@unzip -o $(PROTOC_ZIP) -d /usr/local
	@rm -f $(PROTOC_ZIP)

setup-protobuf-mac: ## install protobuff on mac
	@brew install protobuf
	@brew pin protobuf

setup-dep: ## install dep for dependency management
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@dep version

dep: ## setup all dependencies
	@dep ensure

build: ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)


help:
	@IFS=$$'\n' ; \
    help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
    for help_line in $${help_lines[@]}; do \
        IFS=$$'#' ; \
        help_split=($$help_line) ; \
        help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        printf "%-30s %s\n" $$help_command $$help_info ; \
    done
