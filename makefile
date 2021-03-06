TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'examples' | grep -v '.history')
WIDGETS?=$$(find ./definitions/studio/widgets/* -execdir basename {} .json ';')

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
	cd tools && go mod vendor
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/gordonklaus/ineffassign
	GO111MODULE=off go get -u github.com/gojp/goreportcard/cmd/goreportcard-cli

reportcard:
	@echo "==> running go report card"
	goreportcard-cli

goreportcard-refresh:
	@echo "==> refresh goreportcard checks"
	curl -X POST -F "repo=github.com/RJPearson94/twilio-sdk-go" https://goreportcard.com/checks

generate:
	go generate  ./...

generate-service-api-version:
	@echo "==> regenerating $(SERVICE) $(API_VERSION)"
	rm -rf ./service/$(SERVICE)/$(API_VERSION)
	cd tools && go run ./cli/codegen/service --definition ../definitions/service/$(SERVICE)/$(API_VERSION) --target ../service/$(SERVICE)/$(API_VERSION)
	goimports -w ./service/$(SERVICE)/$(API_VERSION)

generate-studio-widget:
	@echo "==> regenerating $(WIDGET)"
	cd tools && go run ./cli/codegen/studio --definition ../definitions/studio/widgets/$(WIDGET).json --target ../studio/widgets
	goimports -w ./studio/widgets

generate-all-studio-widgets:
	rm -rf ./studio/widgets
	for widget in $(WIDGETS); do \
		make generate-studio-widget WIDGET=$$widget; \
	done

.PHONY: download build test fmt tools generate generate-service-api-version generate-studio-widget generate-all-studio-widgets reportcard goreportcard-refresh