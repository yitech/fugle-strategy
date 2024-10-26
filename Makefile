BINARY := fstrategy

PACKAGE="github.com/yitech/fugle-strategy"
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y-%m-%dT%H:%M:%S')
LDFLAGS := -X '${PACKAGE}/internal/env.Version=${VERSION}' \
           -X '${PACKAGE}/internal/env.CommitHash=${COMMIT_HASH}' \
           -X '${PACKAGE}/internal/env.BuildTime=${BUILD_TIMESTAMP}'

build:
	rm -rf docs
	swag init --parseDependency --parseInternal -g cmd/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o bin/$(BINARY)-linux-amd64 cmd/*.go

clean:
	rm -rf bin/*
