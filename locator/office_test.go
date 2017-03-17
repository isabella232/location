package locator_test

import (
	"net/url"

	. "github.com/thoughtbot/location/locator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Office", func() {
	Describe("URL", func() {
		It("appends the slug to the base URL", func() {
			o := Office{Slug: "shell-less-terrestrial-mollusc"}

			baseURL, err := url.Parse("http://example.com")
			Expect(err).ToNot(HaveOccurred())

			expectedURL := "http://example.com/shell-less-terrestrial-mollusc"
			result := o.URL(*baseURL)

			Expect(result.String()).To(Equal(expectedURL))
		})
	})
})
