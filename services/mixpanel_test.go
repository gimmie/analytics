package services_test

import (
	"encoding/base64"
	"encoding/json"
	. "github.com/llun/analytics/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mixpanel", func() {

	var (
		service     Mixpanel
		mockNetwork MockNetwork
	)

	BeforeEach(func() {
		mockNetwork = MockNetwork{}
		service = Mixpanel{&mockNetwork, "token"}
	})

	Context("#GetName", func() {
		It("should return mixpanel as name", func() {
			name := service.GetName()
			Expect(name).To(Equal("MixPanel"))
		})
	})

	Context("#Send", func() {

		It("should do a get request mixpanel api with base64 encoding data", func() {
			in := Input{
				Event: "view",
				Data: map[string]interface{}{
					"reward": "Nexus5",
				},
			}
			output := service.Send(in)

			Expect(output.Success).To(Equal(true))

			Expect(mockNetwork.Url).To(Equal("https://api.mixpanel.com/track/"))
			expect := map[string]interface{}{
				"event": "view",
				"properties": map[string]interface{}{
					"reward": "Nexus5",
					"token":  "token",
				},
			}
			data, _ := json.Marshal(expect)
			Expect(mockNetwork.Data).To(Equal("data=" + base64.StdEncoding.EncodeToString(data)))

		})

		It("should echo succes=true from Mixpanel.com", func() {
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
					"token":  "token",
				},
			}
			data, _ := json.Marshal(expect)
			Expect(output).To(Equal(data))

		})

	})

})
