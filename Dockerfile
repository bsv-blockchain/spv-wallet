# Get Golang for builder
FROM golang:1.26.3@sha256:313faae491b410a35402c05d35e7518ae99103d957308e940e1ae2cfa0aac29b as builder

# Set the working directory
WORKDIR /go/src/github.com/bsv-blockchain/spv-wallet

COPY . ./

# Build binary
RUN GOOS=linux go build -o spvwallet cmd/main.go

# Get runtime image
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:24650313873554b6ba16c1a1b6b9f9142604f6ab735113e1695faf2dd07fdede

# Version
LABEL version="1.0" name="SPVWallet"

# Set working directory
WORKDIR /

# Copy binary to runner
COPY --from=builder /go/src/github.com/bsv-blockchain/spv-wallet/spvwallet .

# Set entrypoint
ENTRYPOINT ["/spvwallet"]
