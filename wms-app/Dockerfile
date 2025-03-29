# Use Golang base image
FROM golang:1.23.3 AS builder

# Set environment for static build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build a static binary
RUN go build -o wms-app ./cmd

# Use a lightweight Alpine image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary
COPY --from=builder /app/wms-app .

# Ensure the binary is executable
RUN chmod +x wms-app

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./wms-app"]
