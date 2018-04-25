PREFIX := $(GOPATH)
BINDIR := $(PREFIX)/bin
SOURCE := main.go observer/*.go

all: fmt obs-example

obs-example: $(SOURCE)
	go build -o obs-example main.go

.PHONY: fmt
fmt: $(SOURCE)
	gofmt -s -l -w $(SOURCE)

.PHONY: lint
lint: $(SOURCE)
	golint -min_confidence 0.9 observer/...

.PHONY: clean
clean:
	$(RM) obs-example

.PHONY: test-unit
test-unit:
	@echo "running unit tests"
	@go test $(shell go list ./... | grep -v vendor)

.PHONY: install
install: fmt observer
	install -D -m0755 obs-example $(DESTDIR)$(BINDIR)/obs-example

.PHONY: vendor
vendor:
	dep ensure
