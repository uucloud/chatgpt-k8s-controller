# Base image
FROM golang:1.19 AS build

# Set the working directory for the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o controller

# Build the lightweight container
FROM scratch
WORKDIR /app

# Copy the binary file
COPY --from=build /app/controller .

# Run the application
CMD ["./controller"]