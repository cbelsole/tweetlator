FROM golang:1.8-alpine

RUN set -ex \
    && apk add --no-cache ca-certificates glide git bash

WORKDIR /go/src/github.com/cbelsole/tweetlator

COPY . .

RUN GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=0 /go/src/github.com/cbelsole/tweetlator/app .

CMD ["./app"]
