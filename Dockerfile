# syntax=docker/dockerfile:1

FROM golang:latest

# Install vim and nano
# RUN apt-get update && apt-get install -y vim nano

# Set destination for COPY
WORKDIR /app
ADD . /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /cloud-spend

# Expose port 8080 for the app
EXPOSE 8080

# Run the app using 'air'
CMD [ "air" ,"-c",".air.toml"]
