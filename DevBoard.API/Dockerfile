FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/api

FROM alpine:3.19

WORKDIR /app

# Create non-root user
RUN addgroup -g 1001 -S appgroup && adduser -S appuser -u 1001

# Copy only compiled binary from builder
COPY --from=builder --chown=appuser:appgroup /app/app .

# Metadata labels
LABEL maintainer="DevBoard Team"
LABEL description="DevBoard API service"

ENV HOST=0.0.0.0
ENV PORT=8080

EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
CMD test -f /app/app || exit 1

# Run as non-root user
USER appuser

# Security: read-only root filesystem (optional at runtime)
# docker run --read-only --tmpfs /tmp
CMD ["./app"]
