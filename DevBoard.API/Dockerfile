FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/api

FROM alpine:3.19

WORKDIR /app

RUN addgroup -g 1001 -S appgroup && adduser -S appuser -u 1001

COPY --from=builder --chown=appuser:appgroup /app/app .

ENV HOST=0.0.0.0
ENV PORT=8080

EXPOSE 8080

USER appuser

CMD ["./app"]
