GO=go
GOPATH=${HOME}/goprojs

all: build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOPATH=$(GOPATH) $(GO) install ./cmd/echo_server
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOPATH=$(GOPATH) $(GO) install ./cmd/echo_client
	mkdir -p ./bin
	cp -r $(GOPATH)/bin/echo_server ./bin/
	cp -r $(GOPATH)/bin/echo_client ./bin/

clean:
	rm -rf $(GOPATH)/bin
	rm -rf $(GOPATH)/pkg
	rm -rf bin
