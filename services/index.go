package services

import (
	"net/http"
)

type Input struct {
	Event string
	Data  map[string]interface{}
}

type Output struct {
	Success bool
}

type Service interface {
	Send(in Input) Output
	GetName() string
}

type Network interface {
	Request(url string, data string) string
}

type NetworkWrapper struct {
}

func (n NetworkWrapper) Request(url string, data string) string {
	resp, _ := http.Get(url + "?" + data)
	return string(resp.StatusCode)
}
