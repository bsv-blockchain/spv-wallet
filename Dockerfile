# Get Golang for builder
FROM golang:1.27.1@sha256:3680233e3204827fbdc66088528ae6d4b3d034f51d03a99d454f6de034888244 as builder

# Set the working directory
WORKDIR /go/src/github.com/bsv-blockchain/spv-wallet

COPY . ./

# Build binary
RUN GOOS=linux go build -o spvwallet cmd/main.go

# Get runtime image
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:beeada7dd17903dfb69fd5f6916c054720bf28a52daaa2f7a1910a1394244bd2

# Version
LABEL version="1.0" name="SPVWallet"

# Set working directory
WORKDIR /

# Copy binary to runner
COPY --from=builder /go/src/github.com/bsv-blockchain/spv-wallet/spvwallet .

# Set entrypoint
ENTRYPOINT ["/spvwallet"]
