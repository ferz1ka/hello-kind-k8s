# Stage 1: Build stage
FROM golang:1.15 AS build

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2: Final stage
FROM scratch

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/server .

# Set the entrypoint command
ENTRYPOINT ["/app/server"]