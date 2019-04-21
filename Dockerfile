FROM golang:1.12 AS builder
# RUN apt-get update && apt-get install -y git
WORKDIR /go/src/github.com/subhaze/cloudflare-ip-updater
COPY main.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest
RUN apk update \
	&& apk upgrade \
	&& apk add --no-cache \
	ca-certificates \
	&& update-ca-certificates 2>/dev/null || true
WORKDIR /root/
COPY --from=builder /go/src/github.com/subhaze/cloudflare-ip-updater/server .
EXPOSE 8080
ENTRYPOINT ["./server"]