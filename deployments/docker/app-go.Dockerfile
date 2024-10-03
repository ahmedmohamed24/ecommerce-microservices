# Start from the official Golang image to build our application
FROM golang:1.23.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
# Assuming you are using Go Modules (you should be!)
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
CMD ["sh", "-c", "./app migrate > /dev/null 2>&1 && ./app serve"]


