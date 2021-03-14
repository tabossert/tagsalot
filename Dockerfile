# syntax=docker/dockerfile:1.2.1
FROM golang:1.16.0-alpine3.12

WORKDIR /go/src/app

RUN addgroup -S tagsalot && adduser -S -G tagsalot tagsalot

COPY ./src/ ./src
COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN apk add --no-cache git==2.26.2-r0 && \
    go get -d -v ./... && \
    apk del git

RUN cd src/cmd && go build -o /go/src/app/tagsalot && rm -rf src go.mod go.sum
RUN chown -R tagsalot:tagsalot /go/src/app

USER tagsalout

CMD ["./tagsalot"]
