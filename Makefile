

.PHONY: linter-install
linter-install: ## Install linter
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.6


.PHONY: test
test: ## Run all the tests
	go test -covermode=atomic -race -coverprofile=coverage.txt -count=1 ./...


.PHONY: lint
lint: ## Run all the linters
	golangci-lint run


.PHONY: build
build: ## Build a the project
	go build -v ./cmd/mowers