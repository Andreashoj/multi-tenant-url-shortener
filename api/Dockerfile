FROM golang:1.25-alpine

WORKDIR /app

# Install air for hot reload
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Run with air for hot reload
CMD ["air", "-c", ".air.toml"]
