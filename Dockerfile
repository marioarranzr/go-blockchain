FROM golang

ADD . /go/src/github.com/mario800ml/go-blockchain

WORKDIR /go/src/github.com/mario800ml/go-blockchain

RUN go get github.com/davecgh/go-spew/spew
RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv

RUN go install github.com/mario800ml/go-blockchain

ENTRYPOINT /go/bin/go-blockchain

EXPOSE 8080