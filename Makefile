BINARY_NAME=todo
CMD_PATH=./cmd/todo

VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILT_BY ?= $(shell whoami)

LDFLAGS := -s -w \
  -X main.version=$(VERSION) \
  -X main.commit=$(COMMIT) \
  -X main.date=$(DATE) \
  -X main.builtBy=$(BUILT_BY)

build:
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) $(CMD_PATH)

run:
	go run cmd/todo/main.go

clean:
	rm -f todo

lint:
	go fmt ./...

release:
	@echo "Bumping version..."
	git tag v$(VERSION)
	git push origin v$(VERSION)

