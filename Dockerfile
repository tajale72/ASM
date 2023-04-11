# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Compile the Go app
RUN go build -o main .

# Set the command to run the executable
CMD ["/app/main"]