.PHONY: lint
lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.22.2; \
	fi
	golangci-lint run -v

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: install
install:
	go install -ldflags "-X \"main.version=$$(git rev-parse --short=7 @)\"" .
