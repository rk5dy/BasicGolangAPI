FROM golang:1.9.1-alpine3.6

RUN apk --update add git openssh curl make gcc g++ python linux-headers binutils-gold gnupg libstdc++ && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /go/src/app

RUN go get github.com/julienschmidt/httprouter \
  github.com/lib/pq

COPY . .

CMD ["go", "run", "app.go"]
