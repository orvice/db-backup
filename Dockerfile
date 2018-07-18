FROM golang:1.10 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/db-backup
WORKDIR /go/src/github.com/orvice/db-backup
ADD .  /go/src/github.com/orvice/db-backup

RUN go get
RUN CGO_ENABLED=0 go build


FROM alpine:3.7
RUN apk update
RUN apk add --no-cache mysql-client

RUN mkdir -p /app/bin
WORKDIR /app/bin

COPY --from=builder /go/src/github.com/orvice/db-backup/db-backup .

ENTRYPOINT [ "/app/bin/db-backup" ]