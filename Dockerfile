FROM golang:1.22-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o /app/main .

# Stage 2: Create a minimal image to run the application
FROM alpine

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]