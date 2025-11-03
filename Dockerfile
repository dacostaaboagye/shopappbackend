#syntax=docker/dockerfile
FROM golang:1.24-alpine AS builder

#set working directory
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

#copy source code
COPY . .

#build the binaries
RUN go build -o main ./cmd

# --- runtime image ---
FROM alpine:3.19
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
EXPOSE 8080

# Run the binary
CMD ["./main"]