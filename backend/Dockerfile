# ======= STAGE 1: Build the application =======
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go binary (the output binary will be in /app/bin)
RUN go build -o /app/bin/web_app ./cmd

# ===========================================
# ======== STAGE 2: Create lightweight image =======
# ===========================================

FROM alpine:latest

# Set environment variables
ENV GIN_MODE=release

# Create a directory for the app
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/bin/web_app .

# Expose the application port
EXPOSE 8090

# Run the application
CMD ["./web_app"]
