package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	. "github.com/llun/analytics/services"
	"log"
	"net/http"
)

func main() {
	network := NetworkWrapper{}

	mixpanel := Mixpanel{network}
	ga := GA{network}
	aggregator := Aggregator{[]Service{mixpanel, ga}}

	m := martini.Classic()
	m.Post("/send", func(res http.ResponseWriter, req *http.Request) {
		header := res.Header()
		header.Add("Content-Type", "application/json")
		res.WriteHeader(200)

		decoder := json.NewDecoder(req.Body)
		var input map[string]interface{}
		err := decoder.Decode(&input)
		if err != nil {
			panic(-1)
		}
		log.Println(input)

		fmt.Fprintf(res, aggregator.Send(input))
	})
	m.Run()
}
