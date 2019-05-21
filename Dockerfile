FROM golang:1.12.5-alpine3.9

# Install Protocol Buffers v3 & git
RUN apk add --update --no-cache protobuf git

# Install gRPC
RUN go get -u google.golang.org/grpc

# Install the protoc plugin for Go
RUN go get -u github.com/golang/protobuf/protoc-gen-go
