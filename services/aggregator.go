package services

import (
	"encoding/json"
)

type Aggregator struct {
	Services []Service
}

func (a *Aggregator) Send(v map[string]interface{}, ip string) string {
	event, hasName := v["event"].(string)
	data, hasData := v["data"].(map[string]interface{})

	prepare := map[string]interface{}{
		"success":  false,
		"services": map[string]interface{}{},
	}

	if !hasName || !hasData {
		bytes, _ := json.Marshal(prepare)
		return string(bytes)
	}
	in := Input{event, data, ip}
	prepare["success"] = true
	for _, service := range a.Services {
		output := service.Send(in)

		services, _ := prepare["services"].(map[string]interface{})
		services[service.GetName()] = output.Success
	}

	bytes, _ := json.Marshal(prepare)
	return string(bytes)
}
