FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files to the container
COPY metrics_generator.go .
COPY go.mod .
COPY go.sum .

# Download and install Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o metricsgenerator

# Expose the port your application listens on
EXPOSE 8082

# Run the binary built above
CMD ["./metricsgenerator"]
