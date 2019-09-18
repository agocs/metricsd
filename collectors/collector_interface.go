package collectors

import "github.com/agocs/metricsd/structs"
import "github.com/vaughan0/go-ini"

type CollectorInterface interface {
	Enabled() bool
	Report() (structs.MetricSlice, error)
	Setup(ini.File)
	State(bool)
}
