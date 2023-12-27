# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

COPY . /app

# Build the Go application inside the container
RUN go build -o main


# Expose a port if your Go application listens on a specific port
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]
