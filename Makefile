# Variables
BIN_DIR := $(HOME)/.local/bin
TARGET := bluecpprint
SRC := cmd/main.go

# Check for Go at the start
ifndef GO
GO := $(shell command -v go 2>/dev/null)
endif

ifeq ($(GO),)
$(error "Go is not installed or not in PATH. Please install Go to proceed.")
endif

.PHONY: all build run install uninstall clean check

all: build

check:
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go is not installed. Aborting."; exit 1; }

build: check
	@echo "Building..."
	@go build -o $(TARGET) $(SRC) 2>&1

run: check
	@go run $(SRC)

install: build
	@echo "Installing..."
	@mkdir -p $(BIN_DIR)
	@if [ -f $(BIN_DIR)/$(TARGET) ]; then \
	    echo "File $(BIN_DIR)/$(TARGET) already exists."; \
	    read -p "Do you want to replace it? [y/N] " ans; \
	    if [ "$$ans" != "y" ] && [ "$$ans" != "Y" ]; then \
	        echo "Installation aborted."; \
	        exit 1; \
	    fi; \
	    echo "Replacing existing file..."; \
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
