FROM golang:1.9.2-stretch

ENV VERSION=v9.2.0 YARN_VERSION=latest

# For base builds
# ENV CONFIG_FLAGS="--fully-static --without-npm" DEL_PKGS="libstdc++" RM_DIRS=/usr/include

RUN apt-get update && apt-get install -y sudo  \
  && curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash - \
  && apt-get install -y nodejs \
  && npm i -g artillery

WORKDIR /go/src/app

RUN go get github.com/julienschmidt/httprouter \
  github.com/lib/pq

COPY . .

CMD ["go", "run", "app.go"]
