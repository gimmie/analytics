package app

import (
	. "github.com/llun/analytics/services"
)

type Aggregator struct {
	Services []Service
}

func (a *Aggregator) Send(v map[string]interface{}) string {
	return ""
}
