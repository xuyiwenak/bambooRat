FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/user-web

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-web .

FROM alpine:3.9 as prod

WORKDIR /root/user-web

COPY --from=0 /go/src/user-web .

CMD ["./user-srv"]