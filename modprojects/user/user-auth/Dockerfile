FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/user-auth

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-auth .

FROM alpine:3.9 as prod

WORKDIR /root/user-auth

COPY --from=0 /go/src/user-auth  .

CMD ["./user-auth"]