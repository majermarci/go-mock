# go-mock

Simple server to mock responses for any given path and request.

## Config file structure

The configuration is a yaml file containing paths, and each path can have one of each http method defined with a response.
The response on methods must be uppercase and contain a response status. Additionally they can have a body and any number of headers.

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

## Running the mock server

You can run the application in the following ways:

- [Docker Container](#docker-container)
- [Helm Chart](#helm-chart)
- [Locally with Go](#locally-with-go)

### Docker Container

```bash
docker run -p 8080:8080 -v config.yaml:/app/config.yaml ghcr.io/majermarci/go-mock:latest
```

Optionally you can use Docker Compose as well

```yaml
version: '3'

services:

  go-mock:
    image: ghcr.io/majermarci/go-mock:latest
    container_name: go-mock
    volumes:
      - config.yaml:/app/config.yaml
    ports:
      - 8080:8080
    restart: unless-stopped
```

### Helm Chart

WIP

### Locally with Go

You must have Go installed at least on version 1.22.2!

1. Clone the repository

```bash
git clone https://github.com/majermarci/go-mock.git
cd go-mock
```

2. Start the server

```bash
go run ./server
```
