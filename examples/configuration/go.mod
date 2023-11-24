module github.com/liuxd6825/dapr-go-sdk/examples/configuration

go 1.21

toolchain go1.21.4

// Needed to validate SDK changes in CI/CD
replace github.com/liuxd6825/dapr-go-sdk => ../../

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/liuxd6825/dapr-go-sdk v1.9.0
	google.golang.org/grpc v1.57.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
