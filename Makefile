BIN_DIR=bin

all: examples

clean:
	@rm -rf $(BIN_DIR)

examples: requirements
	@go build -o $(BIN_DIR)/cryptonote examples/cryptonote.go
	@go build -o $(BIN_DIR)/litecoin examples/litecoin.go

requirements:
	@mkdir -p $(BIN_DIR)
