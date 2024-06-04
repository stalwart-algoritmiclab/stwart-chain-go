.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: init build

NAME := stwart
APP_NAME := $(NAME)d
VERSION := $(shell echo $(shell git describe --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# TODO
ldflags = \
	-X github.com/cosmos/cosmos-sdk/version.Name=$(NAME) \
	-X github.com/cosmos/cosmos-sdk/version.AppName=$(APP_NAME) \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

init:
	sh ./scripts/init_chain.sh

build:
	go build -o ./$(APP_NAME) -mod=readonly $(BUILD_FLAGS) ./cmd/$(APP_NAME)
