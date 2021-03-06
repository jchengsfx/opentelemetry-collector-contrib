module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver

go 1.14

require (
	github.com/shirou/gopsutil v3.20.12+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/testcontainers/testcontainers-go v0.9.0
	go.opentelemetry.io/collector v0.18.1-0.20210126001456-ef40657447c0
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.16.0
)

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible
