FROM golang:1.16 AS builder
ENV GO111MODULE=off
ENV	CGO_ENABLED=0
ENV	GOOS=linux
ENV	GOARCH=amd64
ENV VERSION=1.0
WORKDIR /build
COPY . .
RUN go build -o goweb .

FROM scratch
COPY --from=builder /build/goweb /
EXPOSE 8080
ENTRYPOINT ["/goweb"]

