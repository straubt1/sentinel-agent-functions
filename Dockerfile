# First stage: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Create a Go web server that will serve our functions as HTTP endpoints
COPY app/*.go .
COPY app/go.* .

# Build the web server
RUN go build -o agentfunctionserver .

# Second stage: Use the TFC agent as base
FROM hashicorp/tfc-agent:latest

# Create hooks, not currently used
RUN mkdir -p /home/tfc-agent/.tfc-agent
ADD --chown=tfc-agent:tfc-agent hooks /home/tfc-agent/.tfc-agent/hooks

# Copy the Go web server from the builder stage
COPY --from=builder /app/agentfunctionserver /usr/local/bin/agentfunctionserver

# Start the web server in the background when container starts
COPY entrypoint-wrapper.sh /entrypoint-wrapper.sh
ENTRYPOINT ["/entrypoint-wrapper.sh"]