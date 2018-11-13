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
		newYorkOffice  Office
		londonOffice   Office
		offices        []Office
		fakeIpResolver *locatorfakes.FakeIpResolverInterface
		locator        OfficeLocator
	)

	BeforeEach(func() {
		newYorkOffice = Office{Slug: "new-york-city", Lat: 40.752547, Long: -73.987005}
		londonOffice = Office{Slug: "london", Lat: 51.519741, Long: -0.099063}
		offices = []Office{newYorkOffice, londonOffice}

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

		locatedOffice, distanceKm, err := locator.Nearest("1.1.1.1")
		Expect(err).ToNot(HaveOccurred())
		Expect(locatedOffice).To(Equal(londonOffice))
		Expect(distanceKm).To(BeNumerically("~", 2.4, 0.2))
	})

	It("returns the New York office when you are near Statue of Liberty", func() {
		statueLibertyLat := 40.6892
		statueLibertyLong := -74.0445
		fakeIpResolver.ResolveCityReturns(statueLibertyLat, statueLibertyLong, nil)

		locatedOffice, distanceKm, err := locator.Nearest("9.9.9.9")
		Expect(err).ToNot(HaveOccurred())
		Expect(locatedOffice).To(Equal(newYorkOffice))
		Expect(distanceKm).To(BeNumerically("~", 8.5, 0.2))
	})

	Context("resolving the IP address fails", func() {
		It("propagates the error", func() {
			expectedError := fmt.Errorf("Oh no")
			fakeIpResolver.ResolveCityReturns(0.0, 0.0, expectedError)

			_, _, err := locator.Nearest("")
			Expect(err).To(MatchError(expectedError))
		})
	})
})
