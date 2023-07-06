package graphite

import (
	"fmt"
)

// Metric is a struct that defines the relevant properties of a graphite metric
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
}

func NewMetric(name, value string, timestamp int64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}

func (metric Metric) String() string {
	return fmt.Sprintf("%s %s %d", metric.Name, metric.Value, metric.Timestamp)
}
