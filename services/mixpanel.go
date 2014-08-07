package services

import (
	"encoding/base64"
	"encoding/json"
)

type Mixpanel struct {
	Network Network
	Token   string
}

func (m *Mixpanel) Parse(in Input) ([]byte, error) {
	in.Data["token"] = m.Token
	prepare := map[string]interface{}{
		"event":      in.Event,
		"properties": in.Data,
	}

	bytes, err := json.Marshal(prepare)
	return bytes, err
}

func (m Mixpanel) GetName() string {
	return "MixPanel"
}

func (m Mixpanel) Send(input Input) Output {
	data, _ := m.Parse(input)
	m.Network.Request("https://api.mixpanel.com/track/", "data="+base64.StdEncoding.EncodeToString(data))
	return Output{Success: true}
}
