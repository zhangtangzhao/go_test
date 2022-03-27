export tag=v1.0
root:
	export ROOT=github.com/zhangtangzhao/service2

build:
	echo "building service2 binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building service2 container"
	docker build -t zhangtangzhao/service2:${tag} .

push: release
	echo "pushing zhangtangzhao/service2"
	docker push zhangtangzhao/service2:${tag}