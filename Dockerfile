FROM golang:1.14.3-alpine AS build

ENV XDG_CACHE_HOME=/tmp/go-build \
    GO111MODULE=on

# First we need to build all development tools
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        curl \
        git \
    # Golang linter
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0

# And finally build our image
FROM golang:1.14.3-alpine

ENV CGO_ENABLED=0 \
    XDG_CACHE_HOME=/tmp/go-build \
    GO111MODULE=on

COPY --from=build \
    /go/bin/golangci-lint \
    /bin/

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
      git \
      make \
      upx
