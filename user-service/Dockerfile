# Use the official Go image as the base image
FROM golang

# Set the working directory inside the container
WORKDIR /userservice

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ../../build/userservice

EXPOSE 10011

# Set the entry point for the container
CMD ["/build/userservice"]
