# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: test build
.DEFAULT_GOAL := help

GIT_COMMIT=$(shell git rev-list -1 HEAD)
GIT_VERSION=$(shell git describe --abbrev=0 --tags)

test: ## run tests
	go test -v -cover -race ./ipsum

build: ## build binaries for distribution
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static" -X main.buildVersion=$(GIT_VERSION) -X main.buildCommitHash=$(GIT_COMMIT)' -o lorem-Linux-x86_64 main.go 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static" -X main.buildVersion=$(GIT_VERSION) -X main.buildCommitHash=$(GIT_COMMIT)' -o lorem-Darwin-x86_64 main.go 
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -ldflags '-s -w  -extldflags "-static" -X main.buildVersion=$(GIT_VERSION) -X main.buildCommitHash=$(GIT_COMMIT)' -o lorem-Linux-armv7l main.go 

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
