
.PHONY: fmt
# go format *.go files
fmt:
	@gofmt -s -w .
	@echo "code formatted"


.PONY: linters-install
linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.41.1; \
	}


.PONY: lint
lint: fmt linters-install
	@golangci-lint run ./...
	@echo "code linted"