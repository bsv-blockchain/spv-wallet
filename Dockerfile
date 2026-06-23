# Get Golang for builder
FROM golang:1.26.4@sha256:792443b89f65105abba56b9bd5e97f680a80074ac62fc844a584212f8c8102c3 as builder

# Set the working directory
WORKDIR /go/src/github.com/bsv-blockchain/spv-wallet

COPY . ./

# Build binary
RUN GOOS=linux go build -o spvwallet cmd/main.go

# Get runtime image
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:850143255ee0d1915f09aaa09f6ed31f24086ba605c323badfbefa95b8c52b0e

# Version
LABEL version="1.0" name="SPVWallet"

# Set working directory
WORKDIR /

# Copy binary to runner
COPY --from=builder /go/src/github.com/bsv-blockchain/spv-wallet/spvwallet .

# Set entrypoint
ENTRYPOINT ["/spvwallet"]
