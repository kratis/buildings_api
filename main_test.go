package main

import (
    "github.com/gorilla/mux"
    "fmt"
    "bytes"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var router *mux.Router

	BeforeEach(func() {
	    router = mux.NewRouter()
	    router.HandleFunc("/buildings", index).Methods("GET")
	    router.HandleFunc("/buildings/{id}", show).Methods("GET")
	    router.HandleFunc("/buildings", create).Methods("POST")
	    router.HandleFunc("/buildings/{id}", delete).Methods("DELETE")
	})

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
				router.ServeHTTP(w, req)
				var buildingsResponse []Building
				body, _ := ioutil.ReadAll(w.Body)
				json.Unmarshal(body, &buildingsResponse)

				Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(len(buildingsResponse)).To(Equal(1))
			})
		})
	})


    Describe("show", func() {
		Context("with valid data", func() {
			BeforeEach(func() {
				buildings = append(buildings, Building{Id: "1", Floors: []int{1,2,3}})
			})

			AfterEach(func() {
				buildings = []Building{}
			})

			It("return correct header", func() {
				req, _ := http.NewRequest("GET", "/buildings/1", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				var buildingsResponse Building
				body, _ := ioutil.ReadAll(w.Body)
				json.Unmarshal(body, &buildingsResponse)

				Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(buildingsResponse.Id).To(Equal("1"))
				Expect(buildingsResponse.Floors).To(Equal([]int{1,2,3}))
			})
		})
	})

    Describe("create", func() {
		Context("with valid data", func() {
			BeforeEach(func() {
				buildings = append(buildings, Building{Id: "1", Floors: []int{1,2,3}})
			})

			AfterEach(func() {
				buildings = []Building{}
			})

            requestBodyString := fmt.Sprintf(`{"address": {"city": "jaipur", country: "india"}, floors: [1,2]}`)
			requestBody := bytes.NewReader([]byte(requestBodyString))
			req, _ := http.NewRequest("POST", "/buildings", requestBody)
			w := httptest.NewRecorder()

			It("return correct header", func() {
				router.ServeHTTP(w, req)
				var buildingsResponse []Building
				body, _ := ioutil.ReadAll(w.Body)
				json.Unmarshal(body, &buildingsResponse)

				Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
			})

			It("return correct response", func() {
				router.ServeHTTP(w, req)
				var buildingsResponse []Building
				body, _ := ioutil.ReadAll(w.Body)
				json.Unmarshal(body, &buildingsResponse)

				Expect(len(buildingsResponse)).To(Equal(2))
				Expect(buildingsResponse[1].Id).To(Equal("2"))
			})
		})
	})	
})