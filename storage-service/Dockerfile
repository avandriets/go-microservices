FROM golang:latest as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o service main.go
FROM scratch
COPY --from=builder /build/service /app/

WORKDIR /app
CMD ["./service"]
