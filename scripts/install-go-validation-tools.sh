#!/usr/bin/env sh

echo "Installing go validation tools ..."
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
echo "Done installing go validation tools."
