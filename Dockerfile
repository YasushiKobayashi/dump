FROM golang:1.10.1-alpine as builder

RUN apk --no-cache add git && \
  go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/YasushiKobayashi/dump
ADD ./ /go/src/github.com/YasushiKobayashi/dump
RUN dep ensure && \
  go build -o dump main.go

FROM alpine:3.7
MAINTAINER Yasushi kobayashi <ptpadan@gmail.com>

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/YasushiKobayashi/dump/dump /dump
ENTRYPOINT ["/dump"]
CMD ["--help"]
