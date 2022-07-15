FROM golang:rc-alpine AS builder

RUN mkdir -p /home/goclient
WORKDIR /home/goclient

ADD . /home/goclient

RUN go mod download
RUN go build -o client

FROM alpine:latest

RUN mkdir -p /home/goclient
WORKDIR /home/goclient
COPY --from=builder /home/goclient/client ./
ENV PORT 8080
EXPOSE 8080

CMD ["./client"]
