FROM golang:1.18.3 as base

FROM base as dev

WORKDIR /go/app
COPY . .



RUN go get -d -v ./...

FROM dev as built
ENV CGO_ENABLED=0
RUN go build -o /usr/bin/app cmd/cli/main.go
