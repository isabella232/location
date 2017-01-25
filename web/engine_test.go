package web_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/thoughtbot/location/web"
	"github.com/thoughtbot/location/web/webfakes"
	"github.com/totherme/nosj"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/totherme/nosj/gnosj"
)

var _ = Describe("Web engine", func() {
	Describe("/v1/nearest", func() {
		var (
			fakeOfficeLocator *webfakes.FakeOfficeLocatorInterface
			request           *http.Request
			recorder          *httptest.ResponseRecorder
			engine            *gin.Engine
		)

		BeforeEach(func() {
			fakeOfficeLocator = &webfakes.FakeOfficeLocatorInterface{}
			request, _ = http.NewRequest("GET", "/v1/nearest", nil)
			recorder = httptest.NewRecorder()
			engine = GetMainEngine(fakeOfficeLocator)
		})

		It("returns the nearest thoughtbot office", func() {
			fakeOfficeLocator.NearestReturns("outer-mongolia", nil)

			clientIP := "127.2.3.4"
			request.Header.Set("X-Forwarded-For", clientIP)

			engine.ServeHTTP(recorder, request)

			Expect(fakeOfficeLocator.NearestCallCount()).To(Equal(1))
			Expect(fakeOfficeLocator.NearestArgsForCall(0)).To(Equal(clientIP))

			Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))

			rawJSON := recorder.Body.String()
			j, err := nosj.ParseJSON(rawJSON)
			Expect(err).NotTo(HaveOccurred())

			slug, _ := j.GetByPointer("/slug")
			Expect(slug).To(BeAString())
			Expect(slug.StringValue()).To(Equal("outer-mongolia"))
		})

		Context("office locator returns an error", func() {
			It("responds 404", func() {
				expectedError := fmt.Errorf("Boom")
				fakeOfficeLocator.NearestReturns("", expectedError)

				engine.ServeHTTP(recorder, request)

				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))

				rawJSON := recorder.Body.String()
				j, err := nosj.ParseJSON(rawJSON)
				Expect(err).NotTo(HaveOccurred())

				errorNode, _ := j.GetByPointer("/error")
				Expect(errorNode).To(BeAString())
				Expect(errorNode.StringValue()).To(Equal(expectedError.Error()))
			})
		})
	})
})
