.PHONY: build-linux build-mac build-windows install clean

BINARY_NAME=jsmodel
SOURCE_DIR=./

build: ## Build for the current system
	go build -o $(BINARY_NAME) $(SOURCE_DIR)

build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 $(SOURCE_DIR)

build-mac: ## Build for macOS
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 $(SOURCE_DIR)

build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe $(SOURCE_DIR)

install: build ## Install on the current system
	go install $(SOURCE_DIR)

clean: ## Clean the built binaries
	rm -f $(BINARY_NAME)*

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
