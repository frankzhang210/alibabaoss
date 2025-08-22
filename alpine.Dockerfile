# Stage 1: Build the Go application
FROM golang:1.23-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app, creating a statically linked binary.
# The -ldflags="-w -s" flag strips debug information, reducing the binary size.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o alibabaoss .

# Stage 2: Create the final, minimal production image
FROM alpine:latest

# Install ca-certificates for HTTPS connections (e.g., to Alibaba Cloud)
RUN apk --no-cache add ca-certificates

# Set run mode to production. This prevents loading dev-only features like swagger.
ENV BEEGO_RUNMODE=prod

WORKDIR /app

# Copy the configuration file needed at runtime.
COPY ./conf/app.conf ./conf/app.conf

# Copy the built binary from the builder stage.
COPY --from=builder /app/alibabaoss .

# Expose the port the app runs on.
EXPOSE 39022

# Command to run the application.
CMD ["./alibabaoss"]

# Example build command:
# sudo docker image build -t alibabaoss-alpine:2.0 -f ./alpine.Dockerfile .
# sudo docker run --rm -p 39022:8080 --name alibabaoss-container alibabaoss-alpine:2.0
