FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy all files from the current directory to the working directory
COPY . .

# Build the Go application
RUN go build -o service1

# Command to run the application
# Only defining the folder name ./service2 because it defaults to main.go
CMD ["./service1"]
