FROM golang:rc-alpine AS builder

RUN mkdir -p /home/goapi
WORKDIR /home/goapi

ADD . /home/goapi

RUN go mod download
RUN go test -v /home/goapi
RUN go build -o api

FROM alpine:latest

RUN mkdir -p /home/goapi
WORKDIR /home/goapi
COPY --from=builder /home/goapi/api ./
ENV PORT 8080
EXPOSE 8080

CMD ["./api"]
