TEST?=$$(go list ./...)

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmt generate
	go install

test: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST) -tags=unit

testacc: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST) -tags=acceptance

fmt:
	@echo "==> Fixing source code with goimports (uses gofmt under the hood)..."
	goimports -w .

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports
	go install github.com/client9/misspell/cmd/misspell

generate:
	go generate  ./...

.PHONY: download build test fmt tools generate