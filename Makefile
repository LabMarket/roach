# Makefile for the Roach project

# The name of the binary to be built
BINARY_NAME=roach

# Get the short git commit hash to use as the build number
BUILD := $(shell git rev-parse --short HEAD)

# Go linker flags to inject the build number into the binary.
# This corresponds to: go build -ldflags "-X 'roach/version.BuildNumber=...'"
LDFLAGS := -ldflags "-X 'roach/version.BuildNumber=$(BUILD)'"

# Default installation path. Can be overridden from the command line.
# e.g., make install INSTALL_PATH=/usr/local/bin
INSTALL_PATH ?= $(HOME)/.local/bin

.PHONY: all build clean install

# The default target executed when you run `make`
all: build install clean

# Build the binary
build:
	@echo "Building $(BINARY_NAME) (version: $(BUILD))..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .
	@echo "$(BINARY_NAME) built successfully."

# Install the binary to the installation path
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@install -d $(INSTALL_PATH)
	@install $(BINARY_NAME) $(INSTALL_PATH)
	@echo "$(BINARY_NAME) installed successfully."

# Clean up the built binary
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
