# BUILD STAGE
FROM golang:1.13 AS builder
WORKDIR /go/app
COPY go.mod go.sum /go/app/
RUN go mod download
COPY . /go/app/
RUN GOOS=linux GOARCH=amd64 CG0_ENABLED=0 go build -a -o piqlit-cli /go/app/cli

# RUN STAGE
FROM postman/newman:ubuntu
WORKDIR /etc/newman
COPY --from=builder /go/app/piqlit-cli /usr/local/bin/
COPY --from=builder /go/app/api/postman/docker-entrypoint.sh /usr/local/bin/
COPY --from=builder /go/app/api/postman/environment.json /etc/newman/
ENTRYPOINT [ "/usr/local/bin/docker-entrypoint.sh" ]