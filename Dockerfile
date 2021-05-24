##############################################################
# Builder image
#
FROM golang:1.16-alpine as builder

RUN apk add --no-cache git gcc g++ libc-dev

WORKDIR /go/src/spamhouse/

COPY go.mod .
COPY go.sum .

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

FROM builder AS build1
# Build the code
# sqlite requires gcc and therefore CGO_ENABLED needs to be set
COPY . .

FROM build1 AS build2
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /usr/local/bin/spamhouse .

##############################################################
# Production image
#
FROM alpine:3.4 AS final

RUN apk add --no-cache ca-certificates curl sqlite && rm -rf /var/cache/apk/*

# Copy the binary from the builder
COPY --from=build2 /usr/local/bin/spamhouse /usr/local/bin/spamhouse
