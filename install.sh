#!/bin/bash

binary_name="go-mock"
architecture=$(uname -m)

# Check if the script is run as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

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

# Download the file and check if it's successful
if ! curl -fsSL "$download_url" -o /usr/local/bin/${binary_name}; then
    echo "Failed to download file"
    exit 1
fi

# Make the file executable and check if it's successful
if ! chmod +x /usr/local/bin/${binary_name}; then
    echo "Failed to make file executable"
    exit 1
fi

echo "${binary_name} installed to /usr/local/bin/${binary_name}"
