FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/registration-relayer
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/registration-relayer /go/src/github.com/rarimo/registration-relayer


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/registration-relayer /usr/local/bin/registration-relayer
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["registration-relayer"]
