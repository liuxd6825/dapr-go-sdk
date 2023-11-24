module github.com/liuxd6825/dapr-go-sdk/examples/hello-world

go 1.21

toolchain go1.21.4

// Needed to validate SDK changes in CI/CD
replace github.com/liuxd6825/dapr-go-sdk => ../../

require (
	github.com/liuxd6825/dapr-go-sdk v0.0.0-00010101000000-000000000000
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
	google.golang.org/grpc v1.57.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
