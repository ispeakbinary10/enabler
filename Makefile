BIN=$(GOPATH)/bin
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

.PHONY: fmt vet build install_deps clean

default: all

all: fmt vet

fmt:
	$(info Checking formatting...)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

install_deps:
	$(info Downloading dependencies...)
	go get -v ./...

vet:
	$(info Vetting...)
	go vet ./...

build: install_deps vet
	$(info Building enabler executable to $(BIN)/enabler)
	go build -o $(BIN)/enabler

clean:
	rm -rf $(BIN)
