# Build stage
FROM golang:1.22.1-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go