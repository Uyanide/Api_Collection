FROM golang:1.24 AS builder

WORKDIR /src

# Cache modules
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org && go mod download

# Copy source and build statically
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/api_collection .

# Final runtime image
FROM alpine:3.19

ARG PUID
ARG PGID
ARG PORT

# Create non-root user
RUN addgroup -g ${PGID} app && adduser -u ${PUID} -G app -S app

# Copy binary
COPY --from=builder /app/api_collection /usr/local/bin/api_collection

USER app
EXPOSE ${PORT}

ENTRYPOINT ["/usr/local/bin/api_collection"]