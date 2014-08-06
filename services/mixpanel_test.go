package services_test

import (
	"encoding/base64"
	"encoding/json"
	. "github.com/llun/analytics/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockNetwork struct {
	Url  string
	Data string
}

func (m MockNetwork) Request(url string, data string) string {
	m.Url = url
	m.Data = data
	return "200"
}

var _ = Describe("Mixpanel", func() {

	var (
		service Mixpanel
		network Network
	)

	BeforeEach(func() {
		network = MockNetwork{}
		service = Mixpanel{network}
	})

	Describe("Mixpanel", func() {

		Context("#GetName", func() {
			It("should return mixpanel as name", func() {
				name := service.GetName()
				Expect(name).To(Equal("MixPanel"))
			})
		})

		Context("#Send", func() {

			It("should request mixpanel api with base64 encoding data", func() {
				in := Input{
					Event: "view",
					Data: map[string]interface{}{
						"reward": "Nexus5",
					},
				}
				output := service.Send(in)

				Expect(output.Success).To(Equal(true))

				mock, _ := network.(MockNetwork)
				Expect(mock.Url).To(Equal("https://api.mixpanel.com/track/"))

				expect := map[string]interface{}{
					"event": "view",
					"properties": map[string]interface{}{
						"reward": "Nexus5",
					},
				}
				data, _ := json.Marshal(expect)
				Expect(mock.Data).To(Equal(base64.StdEncoding.EncodeToString(data)))

			})

		})

		Context("#Parse", func() {

			It("should return JSON string with mixpanel properties", func() {

				in := Input{
					Event: "view",
					Data: map[string]interface{}{
						"reward": "Nexus5",
					},
				}
				output, _ := service.Parse(in)

				expect := map[string]interface{}{
					"event": "view",
					"properties": map[string]interface{}{
						"reward": "Nexus5",
					},
				}
				data, _ := json.Marshal(expect)
				Expect(output).To(Equal(string(data)))

			})

		})

	})

})
