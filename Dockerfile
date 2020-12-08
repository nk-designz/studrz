FROM golang AS builder
COPY srv /go/src/srv
WORKDIR /go/src/srv
RUN go get -d -v . && \
    GOOS=linux GOARCH=amd64 go build -tags static -ldflags="-w -s" -o /studapi
EXPOSE 42069
ENTRYPOINT ["/studapi"]