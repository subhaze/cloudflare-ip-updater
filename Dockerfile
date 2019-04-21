FROM golang:1.12 AS builder
# RUN apt-get update && apt-get install -y git
WORKDIR /go/src/github.com/subhaze/cloudflare-ip-updater
COPY main.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/subhaze/cloudflare-ip-updater/server .
EXPOSE 8080
ENTRYPOINT ["./server"]