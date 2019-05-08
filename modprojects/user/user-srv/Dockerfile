FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/user-srv

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-srv .

FROM alpine:3.9 as prod

WORKDIR /root/user-srv

COPY --from=0 /go/src/user-srv .

CMD ["./user-srv"]