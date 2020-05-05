module github.com/chenlixin93/laracom-go/demo-cli

go 1.14

replace github.com/chenlixin93/laracom-go/demo-service => /Users/abc/go/src/laracom/demo-service

require google.golang.org/grpc v1.29.1

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/chenlixin93/laracom-go/demo-service v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
)
