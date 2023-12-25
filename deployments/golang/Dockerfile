# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install the application dependencies
RUN go mod download

# Copy the rest of the application source code
COPY golang .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]