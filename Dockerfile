# Use official Golang image
FROM golang:latest

# Install Air for hot-reloading
RUN go install github.com/air-verse/air@v1.61.7

# Set working directory
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Expose the application port
EXPOSE 8080

# Command to start Air
CMD ["air", "-c", ".air.toml"]
