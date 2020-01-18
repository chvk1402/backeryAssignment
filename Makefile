GOLINT := golangci-lint
GOFMT := gofmt
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test  ./...

simplify:
	@$(GOFMT) -s -l -w $(SRC)

lint: simplify
	@$(GOLINT) run -v \
	--deadline=5m \
	--disable gochecknoglobals \
	--disable lll \
	--max-same-issues 100

build:
	$(GOBUILD)

test:
	$(GOTEST)
