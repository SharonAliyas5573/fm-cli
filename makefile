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

# Check the operating system
OS := $(shell uname)

ifeq ($(OS),Windows_NT)
	# For Windows
	INSTALL_DIR = C:\bin
	BINARY_NAME = fm.exe
	RM = del
else
	# For Unix-like systems
	RM = rm -f
endif

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME)

# Install the binary
install:
	move $(BINARY_NAME) $(INSTALL_DIR)\$(BINARY_NAME)
	chmod +x $(INSTALL_DIR)/$(BINARY_NAME)

# Clean the project
clean:
	$(GOCLEAN)
	$(RM) $(INSTALL_DIR)\$(BINARY_NAME)

# Run unit tests
test:
	$(GOTEST) -v ./...

# Update dependencies
deps:
	$(GOGET) -u ./...

# Build and install the binary
all: clean build install

.PHONY: build install clean test deps all
