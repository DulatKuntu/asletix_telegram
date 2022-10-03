FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./


# build go app
RUN go mod download
RUN go build -o asletix_telegram ./cmd/main.go
EXPOSE 8444

CMD ["./asletix_telegram"]