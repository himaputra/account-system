# Stage 1: Build the application
FROM golang:1.23.6-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod download

# Argument to specify which service to build
ARG SERVICE

# Copy the rest of the application code
COPY . .

# Build the binary for the specified service
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./apps/${SERVICE}

# Stage 2: Create a minimal runtime image
FROM alpine:latest AS production

# Install Atlas CLI
RUN apk add --no-cache curl && \
    curl -sSf https://atlasgo.sh | sh

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 8080

# Command to run migrations and then start the application
CMD atlas migrate apply --dir "file:///migrations" --url "$DB_URL" && ./app