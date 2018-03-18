FROM golang:1.9 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/db-backup
WORKDIR /go/src/github.com/orvice/db-backup
ADD .  /go/src/github.com/orvice/db-backup

# Download and install any required third party dependencies into the container.
RUN go-wrapper download
# RUN go-wrapper install
RUN CGO_ENABLED=0 go build


FROM mysql

COPY --from=builder /go/src/github.com/orvice/db-backup/db-backup .

ENTRYPOINT [ "./db-backup" ]