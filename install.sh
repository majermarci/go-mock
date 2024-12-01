#!/bin/bash

binary_name="go-mock"
architecture=$(uname -m)

SUDO_PROMPT="Please provide sudo pw for install command:"

# Check if the script is run on Linux
os=$(uname -s)
if [ "$os" != "Linux" ]; then
    echo "This script is only supported on Linux."
    echo "Please build the binary from source, or use go install"
    exit 1
fi

# Determine the download URL for the latest release
if [ "$architecture" = "x86_64" ] || [ "$architecture" = "amd64" ]; then
    download_url="https://github.com/majermarci/${binary_name}/releases/latest/download/${binary_name}-linux-amd64"
elif [ "$architecture" = "aarch64" ] || [ "$architecture" = "arm64" ]; then
    download_url="https://github.com/majermarci/${binary_name}/releases/latest/download/${binary_name}-linux-arm64"
else
    echo "Unsupported architecture, no downloadable release for '$architecture'"
    echo "Please build the binary from source, or use go install"
    exit 1
fi

echo "Downloading binary for $architecture"

# Download the file and check if it's successful
if ! curl -fsSL "$download_url" -o /tmp/${binary_name}; then
    echo "Failed to download file"
    exit 1
fi

# Make the file executable and check if it's successful
if ! sudo install -m 0755 /tmp/${binary_name} /usr/local/bin; then
    echo "Failed to install binary from /tmp!"
    exit 1
else
    echo "${binary_name} installed to /usr/local/bin/${binary_name}"
fi
