# Base image
FROM golang:1.13.8-alpine3.11

# Add token arg
ARG token

# Create build directory
WORKDIR /go/src/github.com/eg-api/demo

# Copy app
COPY . ./

# Environment
ENV API_TOKEN=${token}

# Expose port
EXPOSE 80

# Run
CMD ["go", "run", "main.go"]