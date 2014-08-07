package services_test

import (
	. "github.com/llun/analytics/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	"strings"
)

var _ = Describe("Ga", func() {
	var (
		service     GA
		mockNetwork MockNetwork
	)
	BeforeEach(func() {
		mockNetwork = MockNetwork{}
		service = GA{&mockNetwork}
	})

	Context("#GetName", func() {
		It("should return GA as name", func() {
			name := service.GetName()
			Expect(name).To(Equal("GA"))
		})
	})
	Context("#Send", func() {
		It("should do a request to ga api with post data", func() {
			in := GetMockInput()

			var output Output = service.Send(in)
			Expect(output.Success).To(BeTrue())
			Expect(mockNetwork.Data).ToNot(BeNil())
			Expect(mockNetwork.Url).To(Equal("http://www.google-analytics.com/collect"))
		})
	})

	Context("#FormatGAInput", func() {
		It("should return a GA payload", func() {
			in := GetMockInput()
			var (
				output string
			)
			output = service.FormatGAInput(in)
			data := "v=1&tid=TID&cid=CID&t=event&ec=android&ea=view_reward&el=game_key"
			dataStringArray := strings.Split(data, "&")
			sort.Strings(dataStringArray)
			outputStringArray := strings.Split(output, "&")
			sort.Strings(outputStringArray)

			Expect(outputStringArray).To(Equal(dataStringArray))
		})

	})
})
