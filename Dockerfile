FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd
EXPOSE 8080
CMD ["./main"]