# use official Golang image
FROM golang:1.23.4-alpine3.20

# set working directory
WORKDIR /app

# COPY the source code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

# Expose the port
EXPOSE 8000

# Run the executable
CMD ["./api"]

