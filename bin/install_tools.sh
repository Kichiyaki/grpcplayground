#!/bin/sh

set -o errexit -eo pipefail

cd ./internal/tools

go install \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  github.com/golangci/golangci-lint/cmd/golangci-lint

cd ../..
