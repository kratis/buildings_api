package main

import (
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("index", func() {
		Context("with valid data", func() {
			BeforeEach(func() {
				buildings = append(buildings, Building{Id: "1", Floors: []int{1,2,3}})
			})

			AfterEach(func() {
				buildings = []Building{}
			})

			It("return correct header", func() {
				req, _ := http.NewRequest("GET", "/buildings", nil)
				w := httptest.NewRecorder()
				index(w, req)
				var buildingsResponse []Building
				body, _ := ioutil.ReadAll(w.Body)
				json.Unmarshal(body, &buildingsResponse)

				Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(len(buildingsResponse)).To(Equal(1))
			})
		})
	})
})