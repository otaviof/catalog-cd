APP = catalog-cd
OUTPUT_DIR ?= bin

CMD ?= ./cmd/$(APP)/...
PKG ?= ./pkg/...

BIN = $(OUTPUT_DIR)/$(APP)

GOFLAGS ?= -v
GOFLAGS_TEST ?= -race -cover

ARGS ?=

.EXPORT_ALL_VARIABLES:

.PHONY: $(BIN)
$(BIN):
	go build -o $(BIN) $(CMD) $(ARGS)

default: $(BIN)

.PHONY: run
run:
	go run $(CMD) $(ARGS)