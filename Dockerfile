FROM golang:1.15.2-alpine3.12
WORKDIR /go/src/hello
COPY . .
RUN CGO_ENABLED=0 go get -v ./...
FROM scratch
COPY --from=0 /go/bin/hello /
ENTRYPOINT ["/hello"]
