# Use Golang official image with Debian
FROM golang:1.23.2-alpine


# Set the GOVCS environment variable to disable VCS
ENV GOVCS=off


# Install dependencies for Air (a Go hot-reload tool)

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy all source files into the container
COPY . .

# Install Air for hot-reload during development
RUN go install github.com/air-verse/air@latest


# Verify if .air.toml exists inside the container by listing files
#RUN ls -la /app

# Expose port 8080
EXPOSE 8080

#CMD ["air", "-c", ".air.toml"]

# Run the application
CMD ["./tmp/main"]