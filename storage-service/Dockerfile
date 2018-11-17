FROM golang:latest as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go get -d -v ./...

RUN go build -o service main.go
FROM busybox:glibc
COPY --from=builder /build/service /app/

WORKDIR /app
CMD ["./service"]
