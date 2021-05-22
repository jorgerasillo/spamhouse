##############################################################
# Builder image
#
FROM golang:1.16-alpine as builder

RUN apk add --no-cache git gcc g++ libc-dev

# Copy source files
WORKDIR /go/src/spamhouse/
COPY . .

# Build the code
# sqlite requires gcc and therefore CGO_ENABLED needs to be set
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /usr/local/bin/spamhouse .

##############################################################
# Production image
#
FROM alpine:3.4

RUN apk add --no-cache ca-certificates curl sqlite && rm -rf /var/cache/apk/*

# Copy the binary from the builder
COPY --from=builder /usr/local/bin/spamhouse /usr/local/bin/spamhouse
