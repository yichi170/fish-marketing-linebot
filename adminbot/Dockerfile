FROM golang:rc-alpine AS builder

RUN mkdir -p /home/goadmin
WORKDIR /home/goadmin

ADD . /home/goadmin

RUN go mod download
RUN go build -o admin

FROM alpine:latest

RUN mkdir -p /home/goadmin
WORKDIR /home/goadmin
COPY --from=builder /home/goadmin/admin ./
ENV PORT 8080
EXPOSE 8080

CMD ["./admin"]
