FROM golang as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go get -d -v ./...

RUN go build -o service main.go
FROM busybox:glibc
COPY --from=builder /build/service /app/
COPY --from=builder /build/data /app/data

WORKDIR /app
CMD ["./service"]
