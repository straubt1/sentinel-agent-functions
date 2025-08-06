# First stage: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Create a simple Go web server
COPY app/*.go .
COPY app/go.* .


# Go build
RUN go build -o webserver .

# Second stage: Use the TFC agent as base
FROM hashicorp/tfc-agent:latest

# Copy the Go binary from the builder stage
COPY --from=builder /app/webserver /usr/local/bin/webserver

# Make it executable
# RUN chmod +x /usr/local/bin/webserver

# Start the web server in the background when container starts
COPY entrypoint-wrapper.sh /entrypoint-wrapper.sh

# LABEL org.opencontainers.image.description "Custom Terraform Agent to allow for more flexible workflows in Terraform Enterprise or HCP Terraform."

# RUN mkdir -p /home/tfc-agent/.tfc-agent
# ADD --chown=tfc-agent:tfc-agent hooks /home/tfc-agent/.tfc-agent/hooks

# RUN mkdir -p /home/tfc-agent/mirror
# ADD --chown=tfc-agent:tfc-agent network_mirror /home/tfc-agent/mirror

# Set Hook Timeouts to MAX, https://github.com/hashicorp/tfc-agent/pull/820
ENV TFC_AGENT_HOOK_PRE_PLAN_TIMEOUT=600
ENV TFC_AGENT_HOOK_POST_PLAN_TIMEOUT=600
ENV TFC_AGENT_HOOK_PRE_APPLY_TIMEOUT=600
ENV TFC_AGENT_HOOK_POST_APPLY_TIMEOUT=600

ENTRYPOINT ["/entrypoint-wrapper.sh"]