# go-blockchain Dockerfile used to build docker image on concourse ci

FROM alpine

MAINTAINER Mario Arranz (https://github.com/mario800ml/go-blockchain)

# Put the binary into the container
COPY go-blockchain /

# Add the ailibty to /bin/bash
RUN apk add --update bash && rm -rf /var/cache/apk/*

# Runs the binary when someone uses it.
CMD /go-blockchain
