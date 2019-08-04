PKG := "github.com/delgus/swagger-rest"

all: build

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

test: ## Run unittests
	@go test -short ./... -coverprofile cover.out;
	@go tool cover -func cover.out

gen: ## Get the dependencies
	@swagger generate server -t internal -s web -A library --exclude-main

build: ## Build the binary file
	@go build -a -o library -v $(PKG)/cmd/library

clean: ## Remove previous build
	@rm -f library

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'