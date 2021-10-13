REPO_ROOT := $(shell git rev-parse --show-toplevel)
PRE_PUSH := ${REPO_ROOT}/.git/hooks/pre-push

.DEFAULT_GOAL := help
.PHONY: help
help:  ## display this doc
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup:  ## setup tools
	@./bin/setup.sh

.PHONY: tidy
tidy:
	# tidy
	go mod tidy
	git diff --exit-code

.PHONY: lint
lint: setup ## lint
	# cf. https://golangci-lint.run/usage/linters/
	golangci-lint run --fix --sort-results --exclude-use-default=false --enable asciicheck,bodyclose,cyclop,dupl,durationcheck,errname,errorlint,exhaustive,exportloopref,forbidigo,forcetypeassert,funlen,gci,gochecknoinits,gocognit,goconst,gocritic,gocyclo,goerr113,gofmt,gofumpt,gomnd,goimports,goprintffuncname,gosec,ifshort,importas,makezero,misspell,nakedret,nestif,nilerr,nlreturn,noctx,nolintlint,prealloc,predeclared,promlinter,revive,rowserrcheck,sqlclosecheck,stylecheck,tagliatelle,testpackage,thelper,tparallel,unconvert,unparam,wastedassign,whitespace,wrapcheck,wsl
	git diff --exit-code

.PHONY: test
test:  ## test
	# test
	COLOR=true go test -v -race -p=4 -parallel=8 -timeout=300s -cover -coverprofile=./coverage.txt .
	go tool cover -func=./coverage.txt

.PHONY: ci
ci: setup tidy lint test ## ci
