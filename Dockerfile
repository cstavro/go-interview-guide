# Build stage
FROM golang:1-alpine AS builder

WORKDIR /app

# Copy Go module files and build
COPY server.go .
RUN go build -o server server.go

# Runtime stage
FROM alpine:3.19

ARG USER_ID=1000
ARG GROUP_ID=1000

WORKDIR /app

# Install ca-certificates for HTTPS (needed for external links verification)
RUN apk add --no-cache ca-certificates

# Create a non-root user with the specified UID/GID
RUN addgroup -g ${GROUP_ID} appgroup && \
    adduser -u ${USER_ID} -G appgroup -D appuser

# Copy the built binary
COPY --from=builder /app/server .

# Copy static files and templates
COPY go-interview-guide/ ./go-interview-guide/
COPY templates/ ./templates/

# Create the workspaces directory (volume mount point) and ensure it's owned by the app user
RUN mkdir -p workspaces && chown -R appuser:appgroup /app

# Expose the port the server listens on
EXPOSE 8080

# Run the server as the non-root user
USER appuser

CMD ["./server"]
