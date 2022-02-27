FROM golang:1.16 AS builder
ENV GO111MODULE=on
ENV	CGO_ENABLED=0
ENV	GOOS=linux
ENV	GOARCH=amd64
ENV VERSION=1.0
ENV GOPROXY=https://goproxy.io/
WORKDIR /build
COPY . .
RUN go build -o httpServer .

FROM scratch
COPY --from=builder /build/httpServer /
EXPOSE 8080
ENTRYPOINT ["/httpServer"]
