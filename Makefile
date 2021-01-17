default: all

GO_PACKAGES = $$(go list ./... | grep -v vendor)
GO_FILES = $$(find . -name "*.go" | grep -v vendor | uniq)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o simple-server.linux ./main.go
	mv simple-server.linux docker/tmp

fmt:
	gofmt -s -l -w $(GO_FILES)

all: fmt build-linux
