# Latest golang image on apline linux
FROM golang:1.22.0-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

# Work directory
WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o ./main ./cmd

# Final image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root

# Copy the built application from the builder stage
COPY --from=builder /app .

# Exposing server port
EXPOSE 9090

# Starting our application
# CMD ["go", "run", "cmd/main.go"]
# ENTRYPOINT [ "./app/cmd/main" ]
CMD [ "./main" ]



#////////

# # Use the latest golang image on Alpine Linux as the base image
# FROM golang:1.22.0-alpine AS builder

# # Install necessary dependencies
# RUN apk update && \
#     apk add --no-cache git make bash build-base

# # Set the working directory
# WORKDIR /app

# # Copy and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the application source code
# COPY . .

# # Build the application
# RUN go build -o ./tmp/main ./cmd

# # Install air
# RUN go install github.com/cosmtrek/air@latest

# # Final image
# FROM alpine:latest

# # Install necessary runtime dependencies
# RUN apk --no-cache add ca-certificates

# # Set the working directory
# WORKDIR /app

# # Copy the built application from the builder stage
# COPY --from=builder /go/bin/air /usr/local/bin/air
# COPY --from=builder /app .

# # Expose the port the application runs on
# EXPOSE 9090

# # Run the application with air for live reloading
# CMD ["air"]
