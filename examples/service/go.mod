module github.com/liuxd6825/dapr-go-sdk/examples/service

go 1.23.1

toolchain go1.23.4

// Needed to validate SDK changes in CI/CD
replace github.com/liuxd6825/dapr-go-sdk => ../../

require (
	github.com/liuxd6825/dapr-go-sdk v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.68.1
)

require (
	github.com/dapr/dapr v1.12.1-0.20231030205344-441017b888c5 // indirect
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
