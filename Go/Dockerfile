# Start from base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source from current directory to working directory
COPY . .

# Build the application
RUN cd cmd/app/
RUN go build -o main ./cmd/app

# Expose necessary port
EXPOSE 8080

# Run the created binary executable after wait for mysql container to be up
CMD ["./cmd/app/main"]