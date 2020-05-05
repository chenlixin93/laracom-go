# 创建基于 gRPC/Protobuf 的演示服务接口
#build:
#	protoc -I. --go_out=plugins=grpc:. \
#	  proto/demo/demo.proto

# 通过 Docker 容器化微服务接口（Dockerize）
build:
	protoc -I. --go_out=plugins=grpc:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-service .
run:
	docker run -d -p 9091:9999 -p 8081:8000 laracom-demo-service