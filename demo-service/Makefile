# 创建基于 gRPC/Protobuf 的演示服务接口
#build:
#	protoc -I. --go_out=plugins=grpc:. \
#	  proto/demo/demo.proto

# 通过 Docker 容器化微服务接口（Dockerize）
build:
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-service .
run:
	docker run -d -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns laracom-demo-service
