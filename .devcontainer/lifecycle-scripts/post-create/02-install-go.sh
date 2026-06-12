#!/bin/bash

# Installs the latest stable version of Go system-wide.
# Dependencies: bash, curl, grep, sed, uname, tar, rm, ln, sudo

set -e

echo "=== Installing latest Go ==="

echo "Fetching latest Go version..."
GO_JSON=$(curl -sL https://go.dev/dl/?mode=json)
GO_VERSION=$(printf '%s' "$GO_JSON" | grep '"version"' | head -1 | sed -n 's/.*"version": "go\([^"]*\)".*/\1/p')

if [ -z "$GO_VERSION" ]; then
    echo "Failed to determine latest Go version"
    exit 1
fi

GO_ARCH=$(uname -m)
case "$GO_ARCH" in
    x86_64)  GO_ARCH="amd64" ;;
    aarch64) GO_ARCH="arm64" ;;
    *)       echo "Unsupported architecture: $GO_ARCH"; exit 1 ;;
esac

echo "Installing Go ${GO_VERSION} for linux-${GO_ARCH}..."
curl -sL "https://go.dev/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz" -o /tmp/go.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf /tmp/go.tar.gz
sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go
sudo ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt
rm -f /tmp/go.tar.gz
echo "Go version: $(go version)"

echo "=== Go installation complete ==="
