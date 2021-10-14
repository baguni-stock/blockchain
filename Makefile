PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=blockchain \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=blockchaind \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
	@echo "--> Installing blockchaind"
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
	
cross: go.sum
	@echo "--> Installing cross blockchaind"
	@echo "env GOOS=linux GOARCH=386 go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind"
	@env GOOS=linux GOARCH=386 go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind
	@echo "env GOOS=linux GOARCH=arm go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind"
	@env GOOS=linux GOARCH=arm go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind
	@echo "env GOOS=linux GOARCH=arm64 go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind"
	@env GOOS=linux GOARCH=arm64 go install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchaind
