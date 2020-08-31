module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver

go 1.14

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/awsxray => ../../internal/common/awsxray

require (
	github.com/aws/aws-sdk-go v1.34.13
	github.com/google/uuid v1.1.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/awsxray v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-proto v0.4.0
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.8.1-0.20200824223956-becebeb5b751
	go.uber.org/zap v1.15.0
)
