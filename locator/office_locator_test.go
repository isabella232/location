package locator_test

import (
	"fmt"

	. "github.com/thoughtbot/location/locator"
	"github.com/thoughtbot/location/locator/locatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OfficeLocator#Nearest", func() {
	var (
		offices        []Office
		fakeIpResolver *locatorfakes.FakeIpResolverInterface
		locator        OfficeLocator
	)

	BeforeEach(func() {
		offices = []Office{
			Office{Slug: "new-york", Lat: 40.752547, Long: -73.987005},
			Office{Slug: "london", Lat: 51.519741, Long: -0.099063},
		}

		fakeIpResolver = &locatorfakes.FakeIpResolverInterface{}

		locator = OfficeLocator{
			IPResolver: fakeIpResolver,
			Offices:    offices,
		}
	})

	It("returns the London office when you are near Big Ben", func() {
		bigBenLat := 51.5007
		bigBenLong := -0.116773
		fakeIpResolver.ResolveCityReturns(bigBenLat, bigBenLong, nil)

		slug, err := locator.Nearest("1.1.1.1")
		Expect(err).ToNot(HaveOccurred())
		Expect(slug).To(Equal("london"))
	})

	It("returns the New York office when you are near Statue of Liberty", func() {
		statueLibertyLat := 40.6892
		statueLibertyLong := -74.0445
		fakeIpResolver.ResolveCityReturns(statueLibertyLat, statueLibertyLong, nil)

		slug, err := locator.Nearest("9.9.9.9")
		Expect(err).ToNot(HaveOccurred())
		Expect(slug).To(Equal("new-york"))
	})

	Context("resolving the IP address fails", func() {
		It("propagates the error", func() {
			expectedError := fmt.Errorf("Oh no")
			fakeIpResolver.ResolveCityReturns(0.0, 0.0, expectedError)

			_, err := locator.Nearest("")
			Expect(err).To(MatchError(expectedError))
		})
	})
})
