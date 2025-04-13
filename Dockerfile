# Start with a Golang base image
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code with the new directory structure
COPY cmd/ ./cmd/
COPY internal/ ./internal/
#COPY configs/ ./configs/

# Build the application from the cmd/api directory
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api

# Use a small alpine image for the final image
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Copy config files if needed
#COPY --from=builder /app/configs ./configs/

# Expose the port the API runs on
EXPOSE 6000

# Set environment variable
ENV PORT=6000

# Run the API
CMD ["./api"]
