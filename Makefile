# Makefile for CedarHawk
# Builds static binaries for multiple platforms and packages them for release.

# Define build variables.
BINARY_NAME = cedarhawk
BUILD_DIR = build
RELEASE_DIR = release
VERSION = $(shell git describe --tags --always)

# List of platforms to build for (OS and Arch).
PLATFORMS = \
	"linux amd64" \
	"darwin amd64" \
	"windows amd64"

.PHONY: all clean build release

all: test build

# Run tests.
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts.
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR) $(RELEASE_DIR)

# Build binaries for all defined platforms.
build: clean
	@echo "Building binaries for platforms..."
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		OS=`echo $$platform | cut -d ' ' -f1`; \
		ARCH=`echo $$platform | cut -d ' ' -f2`; \
		output_name=$(BINARY_NAME)-$(OS)-$(ARCH); \
		if [ "$$OS" = "windows" ]; then output_name=$$output_name.exe; fi; \
		echo "Building for $$OS/$$ARCH -> $$output_name"; \
		GOOS=$$OS GOARCH=$$ARCH go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$$output_name ./cmd/main.go; \
	done

# Package the binaries for release.
release: build
	@echo "Packaging release artifacts..."
	@mkdir -p $(RELEASE_DIR)
	@for file in $(BUILD_DIR)/*; do \
		tar -czf $(RELEASE_DIR)/`basename $$file`.tar.gz $$file; \
	done
	@echo "Release artifacts are available in the $(RELEASE_DIR) directory."

