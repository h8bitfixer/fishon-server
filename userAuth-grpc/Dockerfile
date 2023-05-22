# Use the official Go image as the base image
FROM golang

# Set the working directory inside the container
WORKDIR /userAuth-grpc

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ../../build/userAuth-grpc

EXPOSE 50051

# Set the entry point for the container
CMD ["/build/userAuth-grpc"]
