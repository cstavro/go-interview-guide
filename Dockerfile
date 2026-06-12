# Build stage
FROM golang:1-alpine AS builder

WORKDIR /app

# Copy Go module files and build
COPY server.go .
RUN go build -o server server.go

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates for HTTPS (needed for external links verification)
RUN apk add --no-cache ca-certificates

# Copy the built binary
COPY --from=builder /app/server .

# Copy static files and templates
COPY go-interview-guide/ ./go-interview-guide/
COPY templates/ ./templates/

# Create the workspaces directory (volume mount point)
RUN mkdir -p workspaces

# Expose the port the server listens on
EXPOSE 8080

# Run the server
CMD ["./server"]
