FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go build -o /go-docs-generator-action cmd/main.go

ENTRYPOINT ["/go-docs-generator-action"]