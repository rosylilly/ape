package ape_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Route", func() {
	var route *Route

	nilHandler := HandlerFunc(func(req *Request, res *Response) (Marshalable, error) {
		return nil, nil
	})

	BeforeEach(func() {
		route = NewRoute([]string{VerbGet}, "/articles/:id", nilHandler)
	})

	Describe("Route", func() {
		It("should have verbs", func() {
			Expect(route.Verbs).To(Equal([]string{VerbGet}))
		})

		It("should match '/articles/1'", func() {
			Expect(route.Match(VerbGet, "/articles/1")).To(BeTrue())
			Expect(route.Params("/articles/1")).To(Equal(
				map[string]string{
					"id":      "1",
					"_format": "",
				},
			))
		})

		It("should match '/articles/1.json'", func() {
			Expect(route.Match(VerbGet, "/articles/1.json")).To(BeTrue())
			Expect(route.Params("/articles/1.json")).To(Equal(
				map[string]string{
					"id":      "1",
					"_format": "json",
				},
			))
		})
	})
})
