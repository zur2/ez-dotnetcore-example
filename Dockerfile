# Build stage
FROM golang:1.16 AS builder

WORKDIR /app

# Copy the source code
COPY . .

# Run the tests
RUN go test ./...

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/index.html .

# Expose the necessary port
EXPOSE 8080

# Run the application
CMD ["./app"]

