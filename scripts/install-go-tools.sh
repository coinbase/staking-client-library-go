#!/usr/bin/env sh

echo "Installing go tools ..."

go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16.0
go install go.einride.tech/aip/cmd/protoc-gen-go-aip@v0.62.0
go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v0.37.0
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
# We currently use the OpenAPI v2 generator, because protoc-gen-openapiv3 is not yet released.
# We want to use v3 in future as it supports Bearer token authentication which our project makes use of.
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1

echo "Done installing go tools."
