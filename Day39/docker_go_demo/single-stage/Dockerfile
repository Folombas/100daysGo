# Single-stage Docker build for Go application
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV GIN_MODE=release

# Run the application
CMD ["./main", "server"]