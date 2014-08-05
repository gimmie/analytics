package services_test

import (
	"encoding/json"
	. "github.com/llun/analytics/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mixpanel", func() {

	var (
		service Mixpanel
	)

	BeforeEach(func() {
		service = Mixpanel{}
	})

	Describe("Mixpanel", func() {

		Context("#ParseInput", func() {

			It("Should return JSON string with mixpanel properties", func() {

				in := Input{
					Event: "view",
					Data: map[string]interface{}{
						"reward": "Nexus5",
					},
				}
				output := service.Parse(in)

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