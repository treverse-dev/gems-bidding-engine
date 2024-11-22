# Use the official Golang image to create a build artifact.
FROM golang:1.23 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./
# Download Go modules
RUN go mod download

# Copy the local code to the container image.
COPY pkg .

# Build the command inside the container.
# Make sure the path for the output binary is set correctly if needed.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use a Docker multi-stage build to create a lean production image.
FROM alpine:3

# Add ca-certificates in case your application makes external HTTPS calls.
RUN apk add --no-cache ca-certificates

# Copy the built binary from the builder stage.
COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["/server"]
