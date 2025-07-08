# Stage 1: Build the Go binary
FROM golang:1.24.3 AS builder


WORKDIR /app

# Only copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Optional: copy config or env files
COPY cmd/.env .env
 

# Install CA certificates if you make HTTPS requests
RUN apk --no-cache add ca-certificates

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
