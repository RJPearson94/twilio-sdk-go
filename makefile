TEST?=$$(go list ./...)

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmtcheck generate
	go install

test: fmtcheck generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST)

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s .

fmtcheck:
	@echo "==> Checking that code complies with gofmt requirements..."
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

tools:
	@echo "==> installing required tooling..."
	go install github.com/client9/misspell/cmd/misspell

generate:
	go generate  ./...

.PHONY: download build test fmt tools generate