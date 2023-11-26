FROM golang:1.21.3-alpine as builder
# Create and change to the app directory
WORKDIR /app
# Cop go.mod and go.sum if exists
COPY go.* ./
# Download dependencies
RUN go mod download
# Copy the rest of the application source code
COPY . ./
# Build the application
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -o server

# Start a new stage from scratch
FROM gcr.io/distroless/base-debian10
WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server ./server
COPY --from=builder /app/config.json ./config.json

CMD ["/server"]
