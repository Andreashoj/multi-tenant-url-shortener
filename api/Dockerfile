FROM golang:1.25-alpine

WORKDIR /app

# Install air for hot reload and delve for debugging
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Expose port and debug port
EXPOSE 8080 2345

# Run with air for hot reload
CMD ["air", "-c", ".air.toml"]
