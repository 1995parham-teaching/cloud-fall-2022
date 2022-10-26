package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Usage struct {
	SuccessCount prometheus.Counter
	FailedCount  prometheus.Counter
}

func (u Usage) IncSuccess() {
	u.SuccessCount.Add(1)
}

func NewUsage(store string) Usage {
	return Usage{
		SuccessCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "cloud-fall-2022",
			Name:      "database_success_count",
			Help:      "number of success in database access",
			Subsystem: "store",
			ConstLabels: prometheus.Labels{
				"store": store,
			},
		}),
	}
}
