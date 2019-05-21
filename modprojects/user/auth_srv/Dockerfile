FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/auth-srv

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-srv .

FROM alpine:3.9 as prod

WORKDIR /root/auth-srv

COPY --from=0 /go/src/auth-srv .

CMD ["./auth-srv"]