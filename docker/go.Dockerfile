FROM golang:1.23.0
RUN go install github.com/air-verse/air@latest
CMD ["air"]