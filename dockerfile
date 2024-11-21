# Use the official Go image as a base image
FROM golang:1.23.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app ./cmd/main.go

# Expose the application port (adjust as per your application)
EXPOSE 8080

# Command to run the application
CMD ["./app"]
