.PHONY: fmt
fmt:
	@hash goimports > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u golang.org/x/tools/cmd/goimports; \
	fi
	@if [ -n "$$(goimports -l .)" ]; then \
      	echo "Go code is not formatted:"; \
      	goimports -d .; \
      	exit 1; \
    fi;

.PHONY: lint
lint:
	golint ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
