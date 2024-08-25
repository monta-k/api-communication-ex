FROM golang:1.23.0

RUN apt-get update && \
  apt-get install unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip && \
  unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local bin/protoc && \
  unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local include/* && \
  rm -rf protoc-3.9.1-linux-x86_64.zip

RUN go install github.com/99designs/gqlgen@v0.17.49
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1