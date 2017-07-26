FROM golang

ADD . /go/src/github.com/qza/allthingstalkgo

RUN go install github.com/qza/allthingstalkgo

ENTRYPOINT /go/bin/allthingstalkgo

EXPOSE 8080 