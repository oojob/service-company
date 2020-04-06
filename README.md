# Company API
> Company Service is responsible for company management section of job portal microservice

*   Create a company
*   Update a company
*   Delete a company
*   Read a single company by ID
*   Read all company data

## Setup
> Setup to start working on this project

### Install GoLang
[version as of writing: go version go1.12.4 darwin/amd64](https://golang.org/)

### setup `$GOPATH`
```bash
# In your bash profile
export GOPATH="/Users/<user>/<folder>/go"
export PATH=$PATH:$GOPATH/bin
```

> ### IMPORTANT! Make sure this repository is located or clone the project
```bash
# clone the project
cd cd $GOPATH/src/github.com/oojob
git clone @repo
$GOPATH/src/github.com/oojob/service-company
```

### Install protobuf
Mac: `make setup-protobuf-mac`
Linux: `make setup-protobuf-linux`
>   See: [Error](http://google.github.io/proto-lens/installing-protoc.html) if there are any failures

### Setup Go environment

#### Install go dep tool (https://github.com/golang/dep)
```bash
make setup-dep
```

Install go dependencies*

```bash
make setup-go
```
> these need to be managed outside of the vendor/ directory because they are used in proto code generation

## Development
> run the api's locally

### Build Services
```bash
make build
```


## Running company service
> company service is build as an command line application.
> After running `make build` run `./bin/server --help` to view available commands
```bash
company service is responsible for CRUD with company entity
