package locator_test

import (
	"fmt"

	geoip2 "github.com/oschwald/geoip2-golang"
	. "github.com/thoughtbot/location/locator"
	"github.com/thoughtbot/location/locator/locatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IpResolver#ResolveCity", func() {
	It("returns the longitude and latitude of the IP's city", func() {
		db, err := geoip2.Open("../data/GeoLite2-City.mmdb")
		Expect(err).ToNot(HaveOccurred())
		defer db.Close()

		resolver := IpResolver{DB: db}
		long, lat, err := resolver.ResolveCity("37.157.32.218")
		Expect(err).ToNot(HaveOccurred())

		Expect(long).To(Equal(51.5092))
		Expect(lat).To(Equal(-0.0955))
	})

	Context("invalid IP", func() {
		It("returns an error", func() {
			fakeDB := &locatorfakes.FakeIpCityDBInterface{}
			resolver := IpResolver{DB: fakeDB}

			_, _, err := resolver.ResolveCity("foo bar baz")

			Expect(err).To(MatchError("Unable to parse IP"))
			Expect(fakeDB.CityCallCount()).To(Equal(0))
		})
	})

	Context("unresolvable IP", func() {
		It("returns an error", func() {
			expectedError := fmt.Errorf("Boom")

			fakeDB := &locatorfakes.FakeIpCityDBInterface{}
			fakeDB.CityReturns(nil, expectedError)

			resolver := IpResolver{DB: fakeDB}
			_, _, err := resolver.ResolveCity("127.0.0.1")

			Expect(err).To(MatchError("Unable to resolve IP"))
		})
	})
})
