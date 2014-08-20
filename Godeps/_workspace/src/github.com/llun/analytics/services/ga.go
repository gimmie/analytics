package services

import (
	"bytes"
	"fmt"
)

type GA struct {
	Network Network
}

func (i GA) GetName() string {
	return "GA"
}

func (i GA) Send(in Input) Output {
	payload := i.FormatGAInput(in)
	i.Network.Request("http://www.google-analytics.com/collect", payload)
	return Output{true}
}

func (i GA) FormatGAInput(in Input) string {
	prepare := map[string]interface{}{
		"cid": "CID",
		"ea":  in.Event,
		"ec":  "android",
		"el":  "game_key",
		"t":   "event",
		"tid": "TID",
		"v":   "1",
	}

	var buffer bytes.Buffer
	for key, value := range prepare {
		buffer.WriteString(fmt.Sprintf("%s=%s&", key, value))
	}

	var output string = buffer.String()
	return output[0 : len(output)-1]
}
