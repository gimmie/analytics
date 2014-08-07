package services

type MockNetwork struct {
	Url    string
	Data   string
	Status string
}

func (m *MockNetwork) Request(url string, data string) (status int, body string, err error) {

	// url + "?" +  data
	// url: https://api.mixpanel.com/track/
	// data: data=somethingelse
	// will call https://api.mixpanel.com/track/?data=somethingelse
	// and if call is success, return 200

	m.Url = url
	m.Data = data
	return 200, "", nil
}

func GetMockInput() Input {
	in := Input{
		Event: "view_reward",
		Data: map[string]interface{}{
			"reward": "Nexus5",
		},
	}
	return in
}
