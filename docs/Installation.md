# Installation Guide

This document outlines the steps required to install CedarHawk on your system.

## Prerequisites
- **Go:** Ensure you have Go (version 1.16 or later) installed. You can download it from [golang.org](https://golang.org/dl/).
- **Git:** Install Git to clone the repository.

## Steps to Install

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/username/cedarhawk.git
   cd cedarhawk
2. Build the binary:
   ```bash
   go build ./cmd/main.go
   ```
   This will produce a binary named cedarhawk (or cedarhawk.exe on Windows).

3. Run Tests (Optional): Verify the installation by running unit and integration tests:

    go test ./...

## Additional Notes
- If you encounter dependency issues, run:
```bash
go mod tidy
```

- Ensure your Go environment (GOPATH, GOROOT) is correctly configured.

