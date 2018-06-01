.PHONY: build install clean test integration dep release
VERSION=`egrep -o '[0-9]+\.[0-9a-z.\-]+' version.go`
GIT_SHA=`git rev-parse --short HEAD || echo`

build:
	@echo "Building demo..."
	@mkdir -p bin
	@go build -ldflags "-X main.GitSHA=${GIT_SHA}" -o bin/demo .

install:
	@echo "Installing demo..."
	@install -c bin/demo /usr/local/bin/demo

clean:
	@rm -f bin/*

dep:
	@dep ensure
