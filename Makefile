include platform.mk

build-tracker:
	go build -o $(LOCAL_BIN) ./cmd/tracker

build: build-tracker

