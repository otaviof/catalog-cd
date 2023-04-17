APP = catalog-cd
OUTPUT_DIR ?= bin
BIN = $(OUTPUT_DIR)/$(APP)

CMD ?= ./cmd/$(APP)/...
PKG ?= ./pkg/...

GOFLAGS ?= -v
GOFLAGS_TEST ?= -v -cover

ARGS ?=

.EXPORT_ALL_VARIABLES:

.PHONY: $(BIN)
$(BIN):
	go build -o $(BIN) $(CMD) $(ARGS)

default: $(BIN)

.PHONY: run
run:
	go run $(CMD) $(ARGS)

install:
	go install $(CMD)

test: test-unit

.PHONY: test-unit
test-unit:
	go test $(GOFLAGS_TEST) $(CMD) $(PKG) $(ARGS)
