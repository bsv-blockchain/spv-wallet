# Get Golang for builder
FROM golang:1.27.1@sha256:f44f6e88636cfb311f9ebace870ded69d943f227bb3cb27d32ffd84ea18c43ea as builder

# Set the working directory
WORKDIR /go/src/github.com/bsv-blockchain/spv-wallet

COPY . ./

# Build binary
RUN GOOS=linux go build -o spvwallet cmd/main.go

# Get runtime image
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:e5161a7d7d99cf22e4f34b72e111211a399d956d9b0e8714da18e9c4c8151041

# Version
LABEL version="1.0" name="SPVWallet"

# Set working directory
WORKDIR /

# Copy binary to runner
COPY --from=builder /go/src/github.com/bsv-blockchain/spv-wallet/spvwallet .

# Set entrypoint
ENTRYPOINT ["/spvwallet"]
