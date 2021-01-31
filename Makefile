GO=GO111MODULE=on go
GOBUILD=$(GO) build

build:
	$(GOBUILD) ./...
