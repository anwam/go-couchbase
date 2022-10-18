FROM golang:alpine as builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY *.go .
RUN CGO=0 GOOS=linux go build -o go-couchbase .

FROM alpine:latest as app
COPY --from=builder /app/go-couchbase /bin/
EXPOSE 3000
ENTRYPOINT [ "go-couchbase" ]

