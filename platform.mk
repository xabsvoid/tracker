LOCAL_BIN:=$(CURDIR)/bin

# golang

export GO111MODULE=on
export GOBIN:=$(LOCAL_BIN)

GOLANGCI_LINT_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_LINT_VER:=1.63.4
install-golangci-lint:
ifeq ($(wildcard $(GOLANGCI_LINT_BIN)),)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VER)
endif

lint: install-golangci-lint
	$(GOLANGCI_LINT_BIN) run ./...

# api

BUF_BIN:=$(LOCAL_BIN)/buf
BUF_VER:=1.49.0
install-buf:
ifeq ($(wildcard $(BUF_BIN)),)
	go install github.com/bufbuild/buf/cmd/buf@v$(BUF_VER)
endif

api: install-buf
	$(BUF_BIN) generate

api-lint: install-buf
	$(BUF_BIN) lint
