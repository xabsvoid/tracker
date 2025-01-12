LOCAL_BIN:=$(CURDIR)/bin

# golang

export GO111MODULE=on
export GOBIN:=$(LOCAL_BIN)

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
