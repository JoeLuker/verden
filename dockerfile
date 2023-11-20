# Start from the official Go image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY app/ .

# Expose port 8080
EXPOSE 8080

# Run the application using Air
CMD ["air"]
