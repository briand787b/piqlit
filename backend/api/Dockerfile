# BUILD STAGE
FROM golang:1.12 AS builder
WORKDIR /go/app
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -o app ./api

# RUN STAGE
FROM alpine:3.10
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/app/app .
CMD [ "/root/app" ]
