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
}
