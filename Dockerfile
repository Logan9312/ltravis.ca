# Use a multi-stage build to keep the final image light

# Stage 1: Building the Go application
# Start from a lightweight Alpine image with Go installed
FROM golang:alpine AS go-builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go files and other necessary files for the build
COPY . .

# Download Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Stage 2: Setting up Node.js for Tailwind CSS compilation
# Use an Alpine image with Node.js installed
FROM node:alpine AS node-builder

# Set the working directory
WORKDIR /node_app

# Copy the Tailwind CSS file(s)
COPY static/tailwind.css ./static/

# Run the Tailwind CSS compilation
RUN npx tailwindcss -i ./static/tailwind.css -o ./static/output.css

# Stage 3: Create the final lightweight image
# Start from a clean Alpine image
FROM alpine

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory in the container
WORKDIR /app

# Copy the compiled Go binary from the first stage
COPY --from=go-builder /app/main .

# Copy the compiled CSS from the second stage
COPY --from=node-builder /node_app/static/output.css ./static/

# Set the PORT environment variable
ENV PORT 8080

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
