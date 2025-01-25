# Base image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy the Go modules files and download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the application port
EXPOSE 8080

# Set the default command to run Air
CMD ["air","-c",".air.toml"]
