# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = fm

# Installation directory
INSTALL_DIR = /bin

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME)

# Install the binary
install:
	sudo mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	sudo chmod +x $(INSTALL_DIR)/$(BINARY_NAME)

# Clean the project
clean:
	$(GOCLEAN)
	sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)

# Run unit tests
test:
	$(GOTEST) -v ./...

# Update dependencies
deps:
	$(GOGET) -u ./...

# Build and install the binary
all: clean build install

.PHONY: build install install-sudo clean test deps all
