FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git
RUN apk add --update make
RUN apk add --update alpine-sdk

# Set the Current Working Directory inside the container
WORKDIR /tmp/file-sharer-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get github.com/google/wire/cmd/wire

COPY . .

RUN make generate

# Build the Go apps
RUN go build -o ./out/file-sharer-app/file_sharer_api ./cmd/main.go

# Start fresh from a smaller image
FROM alpine:3.9

RUN apk add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app/file-sharer-app

COPY --from=build_base /tmp/file-sharer-app/html/ ./

COPY --from=build_base /tmp/file-sharer-app/out/file-sharer-app/ ./

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["/tmp/file-sharer-app/out/file-sharer-app/"] 