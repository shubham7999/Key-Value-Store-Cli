# Use the official Go image as the base image
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY main.go .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o kvstore

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/kvstore .

# Run the key-value store application when the container starts
CMD ["./kvstore"]
