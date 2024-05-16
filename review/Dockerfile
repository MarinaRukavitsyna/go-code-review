# Use the official Go image as the builder stage
FROM golang:latest AS builder

# Install required dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends git gcc libc-dev

# Set the working directory in the container
WORKDIR /go/src/coupon-service

# Copy the entire source code into the container
COPY . .

# Change the working directory to the export command
WORKDIR /go/src/coupon-service/cmd/export

# Build the Go binary
RUN go build -o main .

# Use a minimal base image for the final stage
FROM alpine:latest

# Copy the binary from the builder stage to the final stage
COPY --from=builder /go/src/coupon-service/cmd/export/main /usr/local/bin/

# Set the entry point for the container
ENTRYPOINT ["main"]

# Expose port 8080
EXPOSE 8080
