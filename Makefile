PREFIX := $(GOPATH)
BINDIR := $(PREFIX)/bin
SOURCE := main.go observer/*.go observer/*/*.go examples/*.go

all: fmt observe

observe: $(SOURCE)
	go build -o observe main.go

.PHONY: fmt
fmt: $(SOURCE)
	gofmt -s -l -w $(SOURCE)

.PHONY: lint
lint: $(SOURCE)
	golint -min_confidence 0.9 observer/...

.PHONY: vet
vet: $(SOURCE)
	go tool vet main.go
	go tool vet observer

.PHONY: clean
clean:
	$(RM) observe

.PHONY: test-unit
test-unit:
	@echo "running unit tests"
	@go test $(shell go list ./... | grep -v examples)

.PHONY: install
install: fmt observe
	install -D -m0755 observe $(DESTDIR)$(BINDIR)/observer

.PHONY: vendor
vendor:
	dep ensure
