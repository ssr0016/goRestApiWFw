FROM golang:1.19.0

# SETUP WORKING DIRECTORY
WORKDIR /usr/src/app

COPY . .
RUN go mod tidy