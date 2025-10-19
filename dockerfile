FROM golang:1.25.2-alpine3.22 AS build

WORKDIR /src
COPY . .

RUN go mod download
RUN go test -v -vet=off -buildvcs=false ./...
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -a -ldflags "-s -w" -installsuffix cgo -o go-mock ./go-mock



FROM alpine:3.22

LABEL org.opencontainers.image.source=https://github.com/majermarci/go-mock
LABEL org.opencontainers.image.description="Go Mock Server"
LABEL org.opencontainers.image.licenses=MIT

USER nobody
WORKDIR /app

COPY --from=build /src/go-mock .
COPY config.yaml .

EXPOSE 8080

ENTRYPOINT [ "/bin/sh", "-c" ]
CMD ["/app/go-mock"]
