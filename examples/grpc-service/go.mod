module github.com/liuxd6825/dapr-go-sdk/examples/grpc-service

go 1.21

toolchain go1.21.4

replace github.com/liuxd6825/dapr-go-sdk => ../../

require (
	github.com/liuxd6825/dapr-go-sdk v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.57.1
	google.golang.org/grpc/examples v0.0.0-20230602173802-c9d3ea567325
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
