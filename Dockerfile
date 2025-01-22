
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make build

# Final stage
FROM alpine:3.18

WORKDIR /app

# Add necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy binary and config
COPY --from=builder /app/bin/walkit .
COPY --from=builder /app/config ./config

# Set environment variables
ENV GIN_MODE=release
ENV APP_ENV=production
ENV POCKET_BASE_URL=http://pocketbase:8090/api

# Expose port
EXPOSE 8080

# Run the application
CMD ["./walkit"]