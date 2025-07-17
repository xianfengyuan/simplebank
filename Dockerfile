# Use the official golang image as a base image
FROM golang:1.24.5-alpine3.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Create a new lightweight image for the final application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./db/migration

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
