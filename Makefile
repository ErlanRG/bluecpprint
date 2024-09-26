# Variables
BIN_DIR = $(HOME)/.local/bin
TARGET = bluecpprint
SRC = cmd/main.go

all: build

build:
	@echo "Building..."
	@go build -o $(TARGET) $(SRC) 2>&1

run:
	@go run $(SRC)

install: build
	@echo "Installing..."
	@mkdir -p $(BIN_DIR)
	@if [ -f $(BIN_DIR)/$(TARGET) ]; then \
	    echo "File $(BIN_DIR)/$(TARGET) already exists. Replacing..."; \
	    rm -f $(BIN_DIR)/$(TARGET); \
	fi
	@cp $(TARGET) $(BIN_DIR)/
	@echo "Installation complete. $(BIN_DIR)/$(TARGET) is now available."

uninstall:
	@echo "Uninstalling..."
	@if [ -f $(BIN_DIR)/$(TARGET) ]; then \
	    rm -f $(BIN_DIR)/$(TARGET); \
	    echo "Uninstallation complete. $(BIN_DIR)/$(TARGET) has been removed."; \
	else \
	    echo "No file to uninstall at $(BIN_DIR)/$(TARGET)."; \
	fi

clean:
	@echo "Cleaning..."
	@rm -f $(TARGET)
	@rm -rf test
	@echo "Clean complete."

.PHONY: all build run install uninstall clean watch
