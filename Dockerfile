FROM golang as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -a -installsuffix -o service main.go
FROM busybox:glibc

WORKDIR /app
CMD ["./service"]
