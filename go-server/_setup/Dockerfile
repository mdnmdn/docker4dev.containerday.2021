FROM golang:1.17.1-alpine as build

RUN apk --update add git && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/* && \
    go get github.com/githubnemo/CompileDaemon github.com/go-delve/delve/cmd/dlv

FROM golang:1.17.1-alpine as app

COPY --from=build /go/bin/CompileDaemon /go/bin/dlv /go/bin/
RUN apk add --no-cache libc6-compat && \
    rm -rf /var/lib/apt/lists/* && \    
    mkdir /app

# ENV HOME /app
ENV CGO_ENABLED 0

WORKDIR /app

#CMD ["sh"]

# docker build  -t golang-compile-daemon:1.17-alpine .
