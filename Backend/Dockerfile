# Use the official Golang base image
FROM golang:1.25

# Set the working directory inside the container
WORKDIR /app

# Copy your Go source code into the container
COPY . .

# Build the Go app
RUN go build -o app main.go

# Run the binary when the container starts
CMD ["./app"]
