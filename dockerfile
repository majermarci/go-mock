FROM golang:1.22.2-alpine3.18 AS build

WORKDIR /src
COPY . .

RUN go mod download
RUN go test -v -vet=off -buildvcs=false ./...
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -a -ldflags "-s -w" -installsuffix cgo -o go-mock .



FROM alpine:3.18

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
