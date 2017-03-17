package locator_test

import (
	"fmt"
	"time"

	. "github.com/thoughtbot/location/locator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OfficeRepo", func() {
	Describe("LoadOffices", func() {
		It("loads the all offices", func() {
			offices, err := OfficeRepo{}.LoadOffices("../data/offices.yaml")
			Expect(err).ToNot(HaveOccurred())

			Expect(offices).To(HaveLen(7))

			Expect(offices[0].Slug).To(Equal("austin"))
			Expect(offices[0].Name).To(Equal("Austin"))
			Expect(offices[0].Lat).To(Equal(30.268592))
			Expect(offices[0].Long).To(Equal(-97.743192))
		})

		Context("unable to load file", func() {
			It("returns an error", func() {
				path := "/tmp/does-not-exist-" + string(time.Now().Unix())
				_, err := OfficeRepo{}.LoadOffices(path)

				expectedError := fmt.Sprintf("Unable to load offices from: %s", path)
				Expect(err).To(MatchError(expectedError))
			})
		})
	})
})
