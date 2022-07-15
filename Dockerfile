FROM golang:1.18.3 as base

FROM base as built

WORKDIR /go/app
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /usr/bin/app cmd/cli/main.go
