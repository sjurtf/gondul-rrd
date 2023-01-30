FROM golang:1.19.5-bullseye as builder

RUN apt update && \
    apt install -y librrd-dev rrdtool

ENV GO111MODULE=on \
    CGO_ENABLED=1
#    GOOS=linux \
#    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app
RUN go build -o binary .

FROM alpine:latest

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=builder /app/binary .

USER app
ENTRYPOINT ["./binary"]