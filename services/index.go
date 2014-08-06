package services

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
