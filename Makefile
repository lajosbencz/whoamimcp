IMAGE_NAME ?= whoamimcp
INSPECTOR_PORT ?= 6274
INSPECTOR_PROXY_PORT ?= 6277

whoamimcp:
	CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o whoamimcp ./cmd/whoamimcp

.PHONY: default
default: check test build

.PHONY: run
run: build
	WHOAMI_PORT_NUMBER=12080 ./whoamimcp

.PHONY: build
build: whoamimcp

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: check
check:
	golangci-lint run

.PHONY: image
image:
	docker build -t $(IMAGE_NAME) .

.PHONY: inspector
inspector:
	docker run -it --rm \
	--name dev-whoamimcp-inspector \
	-e "HOST=0.0.0.0" \
	-p "$(INSPECTOR_PORT):6274" \
	-p "$(INSPECTOR_PROXY_PORT):6277" \
	ghcr.io/modelcontextprotocol/inspector:latest
