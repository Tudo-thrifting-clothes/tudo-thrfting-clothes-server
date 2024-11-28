# Step 1: Build the Go binary
FROM golang:1.22.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules file
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Run the Go binary in a smaller container
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder container
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
