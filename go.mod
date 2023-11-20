module github.com/liuxd6825/dapr-go-sdk

go 1.20

require (
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.3.1
	github.com/liuxd6825/dapr v1.12.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.57.1
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/liuxd6825/components-contrib => gitee.com/liuxd6825/components-contrib v1.7.1-1.0-alaph2
replace github.com/liuxd6825/dapr-components-contrib => ../dapr-components-contrib

replace github.com/liuxd6825/dapr => ../dapr
