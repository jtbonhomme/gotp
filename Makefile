.PHONY: help deps
IMAGES_TAG = ${shell git describe --tags --match '[0-9]*\.[0-9]*\.[0-9]*' 2> /dev/null || echo 'latest'}
GIT_SHA1:=$(shell git rev-parse --short HEAD)

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

lint: ## Go lint the repo.
	@which golangci-lint || (go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.36.0)
	golangci-lint -v --deadline 100s --skip-dirs docs run ./...

test: lint ## Go test the repo.
	go test ./... -cover -coverprofile coverage.out

badge: lint ## Generate a coverage badge.
	which gopherbadger || (go get github.com/jpoles1/gopherbadger)
	gopherbadger

cover: test ## Measure the test coverage.
	which gocov || (go get -u github.com/axw/gocov/gocov)
	which gocov-xml || (go get -u github.com/AlekSi/gocov-xml)
	which gocov-html || (go get -u github.com/matm/gocov-html)
	gocov convert coverage.out | gocov-xml > cover.xml
	gocov convert coverage.out | gocov-html > cover.html
	open cover.html
