#!/usr/bin/make -f

# COMMIT := $(shell git log -1 --format='%H')
NETWORK_TYPE := mmx-v1

export GO111MODULE = on
export GOPROXY = https://goproxy.io
# windows, linux, darwin
export GOOS = linux
export GOARCH = amd64

# include ledger support
include Makefile.ledger

ldflags = -X 'mmx.com/go-mmx/version.NetworkType=$(NETWORK_TYPE)' \
          -X 'mmx.com/go-mmx/version.BuildTags=netgo'

BUILD_FLAGS := -tags "$(build_tags)" -ldflags "$(ldflags)"

all: build

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/windows/mmxd.exe ./cmd/mmxd
	go build -mod=readonly $(BUILD_FLAGS) -o build/windows/mmxcli.exe ./cmd/mmxcli
else
    UNAME_S := $(shell uname -s)
    goos_flag := linux
    ifeq ($(UNAME_S), Darwin)
        goos_flag = darwin
    endif
    ifeq ($(UNAME_S), OpenBSD)
        goos_flag = openbsd
    endif
    ifeq ($(UNAME_S), FreeBSD)
        goos_flag = freebsd
    endif
    ifeq ($(UNAME_S), NetBSD)
        goos_flag = netbsd
    endif
	GOOS=$(goos_flag) GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) -o build/$(goos_flag)/mmxd ./cmd/mmxd
	GOOS=$(goos_flag) GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) -o build/$(goos_flag)/mmxcli ./cmd/mmxcli
endif

build-windows: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o build/windows/mmxd.exe ./cmd/mmxd
	go build -mod=readonly $(BUILD_FLAGS) -o build/windows/mmxcli.exe ./cmd/mmxcli

build-linux: go.sum
# LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build
	go build -mod=readonly -tags "netgo" -ldflags "$(ldflags)" -o build/linux/mmxd ./cmd/mmxd
	go build -mod=readonly -tags "netgo" -ldflags "$(ldflags)" -o build/linux/mmxcli ./cmd/mmxcli

build-mac: go.sum
# LEDGER_ENABLED=false GOOS=darwin GOARCH=amd64 $(MAKE) build
	go build -mod=readonly -tags "netgo" -ldflags "$(ldflags)" -o build/mac/mmxd ./cmd/mmxd
	go build -mod=readonly -tags "netgo" -ldflags "$(ldflags)" -o build/mac/mmxcli ./cmd/mmxcli

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependecies have not been modified"
	@go mod verify
	@go mod tidy

clean:
	rm -rf build/

