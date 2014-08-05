package services

import "encoding/json"

type Mixpanel struct {
}

func (m *Mixpanel) Parse(in Input) (string, error) {
	prepare := map[string]interface{}{
		"event":      in.Event,
		"properties": in.Data,
	}

	bytes, err := json.Marshal(prepare)
	return string(bytes), err
}
