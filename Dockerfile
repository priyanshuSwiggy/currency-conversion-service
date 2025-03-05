# Use the official Golang image as the base image
FROM golang:1.23-alpine
LABEL authors="priyanshu"

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download && go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o server ./server
RUN go build -o client ./client

# Expose the ports for HTTP and gRPC servers
EXPOSE 8085
EXPOSE 50051

# Set environment variables for AWS credentials
ENV AWS_REGION=us-west-2
ENV AWS_ACCESS_KEY_ID=dummy
ENV AWS_SECRET_ACCESS_KEY=dummy

# Command to run the server
CMD ["./server"]
