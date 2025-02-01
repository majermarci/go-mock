# go-mock

[![Go Report Card](https://goreportcard.com/badge/github.com/majermarci/go-mock)](https://goreportcard.com/report/github.com/majermarci/go-mock)
[![Go Reference](https://pkg.go.dev/badge/github.com/majermarci/go-mock.svg)](https://pkg.go.dev/github.com/majermarci/go-mock)
![License](https://img.shields.io/github/license/majermarci/go-mock?label=License)
[![Docker Status](https://github.com/majermarci/go-mock/actions/workflows/docker.yaml/badge.svg)](https://github.com/majermarci/go-mock/actions/workflows/docker.yaml)
[![Build Status](https://github.com/majermarci/go-mock/actions/workflows/build.yaml/badge.svg)](https://github.com/majermarci/go-mock/actions/workflows/build.yaml)
[![Test Status](https://github.com/majermarci/go-mock/actions/workflows/audit.yaml/badge.svg)](https://github.com/majermarci/go-mock/actions/workflows/audit.yaml)
[![Latest Release)](https://img.shields.io/github/v/release/majermarci/go-mock?logo=github)](https://github.com/majermarci/go-mock/releases)

Simple server to mock responses for any given path and request.

## Configuration Options

### Server Configuration file

The configuration is a YAML file that defines paths, with each path supporting one or more HTTP methods, each of which specifies a response. Responses must include a status code and may optionally include a body and headers.

> [!WARNING]
> By default the configuration file's name must be `config.yaml` and must be in the current working directory when the server is started.
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

### Pre-defined paths

There are curretnly two paths that are reserved:

- `/healthz`: Shows if the server is up and it's version. Requests are not logged here due to liveness probe checks.
- `/paths`: A list of all available endpoints and the methods that can be used on them. Requires admin credentials which are `admin:admin` by default. It can be changed by setting the `MOCK_ADMIN_PASS` environment variable.

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

To quickly start the server using Docker, provide a [configuration file](https://github.com/majermarci/go-mock/blob/main/config.yaml) to the container by mounting it to `/app/config.yaml`. If no configuration file is provided, the server will use the default one from the repository.

#### Using Docker Compose

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

#### Using Docker CLI

```bash
docker run -p 8080:8080 -v ./config.yaml:/app/config.yaml ghcr.io/majermarci/go-mock:latest
```

---

### Helm Chart

With Helm you can customize the config in your own [values file](https://github.com/majermarci/go-mock/blob/main/helm/go-mock/values.yaml) using the `mockConfig` field.
To install the `go-mock` Helm chart from this repository, run the following commands:

```bash
# Add the GitHub repository as a Helm chart repository
helm repo add go-mock https://majermarci.github.io/go-mock

# Update your local Helm chart repositories
helm repo update

# Install the chart
helm upgrade -i my-go-mock go-mock/go-mock -n go-mock --create-namespace
```

---

### Locally with Go

You must have Go installed at least on version `1.22.2`! Then either run the install command:

```bash
go install github.com/majermarci/go-mock/server
```

Or run the server directly from the code:

```bash
# 1. Clone the repository
git clone https://github.com/majermarci/go-mock.git && cd go-mock

# 2. Start the server
go run ./server
```

---

## Binary Installation

You can also download the binary from the [releases page](github.com/majermarci/go-mock/releases) and run it directly.

Another option is to use the following script to download and install the binary to your `/usr/local/bin` directory:

```bash
curl -fsSL https://raw.githubusercontent.com/majermarci/go-mock/main/install.sh | bash
```

After installing you can run it in any directory that has a valid config file.
