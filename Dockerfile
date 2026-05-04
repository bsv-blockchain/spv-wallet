# Get Golang for builder
FROM golang:1.26.2@sha256:b54cbf583d390341599d7bcbc062425c081105cc5ef6d170ced98ef9d047c716 as builder

# Set the working directory
WORKDIR /go/src/github.com/bsv-blockchain/spv-wallet

COPY . ./

# Build binary
RUN GOOS=linux go build -o spvwallet cmd/main.go

# Get runtime image
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.7@sha256:8d0a8fb39ec907e8ca62cdd24b62a63ca49a30fe465798a360741fde58437a23

# Version
LABEL version="1.0" name="SPVWallet"

# Set working directory
WORKDIR /

# Copy binary to runner
COPY --from=builder /go/src/github.com/bsv-blockchain/spv-wallet/spvwallet .

# Set entrypoint
ENTRYPOINT ["/spvwallet"]
