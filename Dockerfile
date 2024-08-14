FROM golang:1.22-alpine as buildbase

ARG CI_JOB_TOKEN

RUN apk add git build-base ca-certificates
WORKDIR /go/src/github.com/rarimo/registration-relayer
COPY . .

RUN git config --global url."https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com".insteadOf https://gitlab.com
RUN git config --global url."https://${CI_JOB_TOKEN}@github.com/".insteadOf https://github.com/
RUN go env -w GOPRIVATE=github.com/*,gitlab.com/*

RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build  -o /usr/local/bin/registration-relayer /go/src/github.com/rarimo/registration-relayer

FROM scratch
COPY --from=alpine:3.9 /bin/sh /bin/sh
COPY --from=alpine:3.9 /usr /usr
COPY --from=alpine:3.9 /lib /lib

COPY --from=buildbase /usr/local/bin/registration-relayer /usr/local/bin/registration-relayer
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["registration-relayer"]
