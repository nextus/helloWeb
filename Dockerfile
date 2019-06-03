FROM golang:latest as builder
ADD . /go/src/helloWeb
WORKDIR /go/src/helloWeb
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install .

FROM alpine:latest
WORKDIR /helloWeb
COPY --from=builder /go/bin/helloWeb .
CMD ["./helloWeb"]
