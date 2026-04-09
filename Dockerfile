FROM golang:1.26.2

WORKDIR /app

COPY . .

EXPOSE 5000
RUN go test ./...
CMD go run ./cmd/main.go