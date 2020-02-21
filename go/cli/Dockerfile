# BUILD STAGE
FROM golang:1.13 AS builder
WORKDIR /go/app
COPY go.mod go.sum /go/app/
RUN go mod download
COPY . /go/app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -o app /go/app/cli

# RUN STAGE
FROM alpine:3.10
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/app/app .
CMD [ "/root/app" ]