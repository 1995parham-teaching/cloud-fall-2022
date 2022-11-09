package config

import (
	"github.com/1995parham-teaching/cloud-fall-2022/internal/metric"
)

func Default() Config {
	return Config{
		Metric: metric.Config{
			Address: ":9090",
			Enabled: true,
		},
	}
}
