FROM golang:1.17

WORKDIR /go/src/gitlab.com/tokend/blob-serv

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/blob-serv gitlab.com/tokend/blob-serv


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/blob-serv /usr/local/bin/blob-serv
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blob-serv"]
