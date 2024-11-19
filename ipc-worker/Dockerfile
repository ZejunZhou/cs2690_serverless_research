# Base image with Go
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy source files
COPY . .

# Build the Go application
RUN go build -o app .

# Create the runtime image
FROM debian:bullseye

# Set environment variables for IronFunctions
ENV FUNC_HANDLER app
ENV FUNC_PORT 8080

# Install runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Create required directories
RUN mkdir -p /tmp/fifo /tmp/shm

# Copy the compiled binary
COPY --from=builder /app/app /app/app

# Expose the function port
EXPOSE 8080

# Start the function
CMD ["/app/app"]
