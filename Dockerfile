# Use Golang official image with Debian
FROM golang:1.23.2-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Air for hot-reload during development
RUN go install github.com/air-verse/air@latest

# Copy the Go modules files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy all source files into the container
COPY . .

# Expose the ports
EXPOSE 8080

#CMD ["air", "-c", ".air.toml"]
