# Use Go 1.24.3 base image
FROM golang:1.24.3-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the full project
COPY . .

# Build the binary
RUN go build -o main .

# Confirm binary was built (debug step)
RUN ls -la /app

# Expose the port your app uses
EXPOSE 5000

# Run the binary
CMD ["/app/main"]
