# Use the official Golang Alpine image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to cache dependencies
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set environment variable for tests
ENV TEST_ENV=true

# Run the tests
CMD ["sh", "-c", "go test ./tests"]
