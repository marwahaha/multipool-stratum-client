BIN_DIR=bin

all: displayer

clean:
	@rm -rf $(BIN_DIR)

displayer: requirements
	@go build -o $(BIN_DIR)/displayer cmd/displayer.go

requirements:
	@mkdir -p $(BIN_DIR)
