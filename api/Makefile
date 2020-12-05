BINARY_FILE=./bin/domane
DIRECTORIES=go list -f '{{.Dir}}' ./... | grep -v /vendor/
PACKAGES=go list ./... | grep -v /vendor/

.PHONY: build
build: clean
	go build -o ${BINARY_FILE} main.go

.PHONY: checks
checks: format vet

.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_FILE}

.PHONY: deps
deps:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: format
format:
	gofmt -w `$(call DIRECTORIES)`

.PHONY: start
start:
	go run main.go start

.PHONY: test
test:
	go test -a -mod=vendor `$(call PACKAGES)`

.PHONY: vet
vet:
	go vet -mod=vendor `$(call PACKAGES)`