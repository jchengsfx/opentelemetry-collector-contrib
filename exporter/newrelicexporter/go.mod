module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/newrelic/newrelic-telemetry-sdk-go v0.8.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.18.1-0.20210126001456-ef40657447c0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
	google.golang.org/protobuf v1.25.0
)
