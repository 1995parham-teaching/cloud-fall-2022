package config

import (
	"github.com/1995parham-teaching/cloud-fall-2022/internal/metric"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/tracing"
)

func Default() Config {
	return Config{
		Metric: metric.Config{
			Address: ":9090",
			Enabled: true,
		},
		Tracing: tracing.Config{
			Enabled: true,
			Agent: tracing.Agent{
				Host: "127.0.0.1",
				Port: "1234",
			},
		},
	}
}
