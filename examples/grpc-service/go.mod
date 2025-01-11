module github.com/liuxd6825/dapr-go-sdk/examples/grpc-service

go 1.23.1

toolchain go1.23.4

replace github.com/liuxd6825/dapr-go-sdk => ../../

require (
	github.com/liuxd6825/dapr-go-sdk v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.68.1
	google.golang.org/grpc/examples v0.0.0-20230602173802-c9d3ea567325
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
