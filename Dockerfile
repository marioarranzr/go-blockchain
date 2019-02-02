FROM golang

ADD . /go/src/github.com/marioarranzr/go-blockchain

WORKDIR /go/src/github.com/marioarranzr/go-blockchain

RUN go get github.com/davecgh/go-spew/spew
RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv

RUN go install github.com/marioarranzr/go-blockchain

ENTRYPOINT /go/bin/go-blockchain

EXPOSE 8080