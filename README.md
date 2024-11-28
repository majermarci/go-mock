# go-mock

[![Go Report Card](https://goreportcard.com/badge/github.com/majermarci/go-mock)](https://goreportcard.com/report/github.com/majermarci/go-mock)
[![Go Reference](https://pkg.go.dev/badge/github.com/majermarci/go-mock.svg)](https://pkg.go.dev/github.com/majermarci/go-mock)
![License](https://img.shields.io/github/license/majermarci/go-mock?label=License)
[![Build Status](https://github.com/majermarci/go-mock/actions/workflows/build.yaml/badge.svg)](https://github.com/majermarci/go-mock/actions/workflows/build.yaml)
[![Test Status](https://github.com/majermarci/go-mock/actions/workflows/audit.yaml/badge.svg)](https://github.com/majermarci/go-mock/actions/workflows/audit.yaml)
[![Latest Release)](https://img.shields.io/github/v/release/majermarci/go-mock?logo=github)](https://github.com/majermarci/go-mock/releases)

Simple server to mock responses for any given path and request.

## Configuration Options

### Server Configuration file

The configuration is a yaml file containing paths, and each path can have one of each http method defined with a response.
The response on methods must be uppercase and contain a response status. Additionally they can have a body and any number of headers.

> [!WARNING]
> By default the config file's name must be `config.yaml` and must be in the current working directory when the server is started.
> If you want to use a different file, you can provide it with the `-c` flag followed by the path to the file.

Format is the following:

```yaml
/<path>:
    <method>:
        status: <code> # Must be an integer!
        body: <response body>
        headers:
            <key>: <value>
```

Example:

```yaml
/book:
    GET:
        status: 201 # Created
        body: |
            {
                "book": {
                    "title": "Example",
                    "state": "Created"
                }
            }
        headers:
            Content-Type: application/json; charset=utf-8
```

---

### Using a different port

By default the server listens on port 8080, but you can change it by setting the `-p` flag folowed by a port number.
Simply run the server with `go-mock -p <port-number>` to change it.

For Docker the easiest way is to use the `-p` flag or the ports section in compose file, so you can map it to a different one.

## Running the mock server

You can run the application in the following ways:

- [Docker Container](#docker-container)
- [Helm Chart](#helm-chart)
- [Locally with Go](#locally-with-go)

---

### Docker Container

For a quick start you can run the server in a Docker container. You must provide a [config file](https://github.com/majermarci/go-mock/blob/main/config.yaml) to the container, and you can do so by mounting it to the `/app/config.yaml` path. Otherwise the server will start with the default config from the repository.

```bash
docker run -p 8080:8080 -v ./config.yaml:/app/config.yaml ghcr.io/majermarci/go-mock:latest
```

Optionally you can use Docker Compose as well:

```yaml
version: '3'

services:

  go-mock:
    image: ghcr.io/majermarci/go-mock:latest
    container_name: go-mock
    volumes:
      - ./config.yaml:/app/config.yaml
    ports:
      - 8080:8080
    restart: unless-stopped
```

---

### Helm Chart

With Helm you can customize the config in your own [values file](https://github.com/majermarci/go-mock/blob/main/helm/go-mock/values.yaml) using the `mockConfig` field.
To install the `go-mock` Helm chart from this repository, run the following commands:

```bash
# Add the GitHub repository as a Helm chart repository
helm repo add go-mock https://majermarci.github.io/go-mock/helm

# Update your local Helm chart repositories
helm repo update

# Install the chart
helm upgrade -i my-go-mock go-mock/go-mock -n go-mock --create-namespace
```

---

### Locally with Go

You must have Go installed at least on version 1.22.2! Then either run the install command:

```bash
go install github.com/majermarci/go-mock/server
```

Or run the server directly:

```bash
# 1. Clone the repository
git clone https://github.com/majermarci/go-mock.git
cd go-mock

# 2. Start the server
go run ./server
```

---

## Binary Installation

You can also download the binary from the [releases page](github.com/majermarci/go-mock/releases) and run it directly.

Another option is to use the following script to download and install the binary to your `/usr/local/bin` directory:

```bash
curl -fsSL https://raw.githubusercontent.com/majermarci/go-mock/main/install.sh | sudo bash
```

After installing you can run it in any directory that has a valid config file.
