FROM golang:1.23.0-alpine

RUN set -eux && \
  apk --update add --no-cache alpine-sdk tzdata&& \
  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0