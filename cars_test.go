package main_test

import (
	"bytes"
	fern "github.com/guidewire/fern-ginkgo-client/pkg/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
)

var _ = ReportAfterSuite("Report to Fern 1", func(report Report) {
	f := fern.New("Sample report 1",
		fern.WithBaseURL("http://localhost:8080/"),
	)

	err := f.Report("Sample report 1", report)

	Expect(err).To(BeNil(), "Unable to create reporter file")
})

//var _ = ReportAfterSuite("Report to Fern 2", func(report Report) {
//	f := fern.New("Sample report 2",
//		fern.WithBaseURL("http://localhost:8080/"),
//	)
//
//	err := f.Report("Sample report 2", report)
//
//	Expect(err).To(BeNil(), "Unable to create reporter file")
//})

var _ = Describe("Car service test", func() {
	Context("When successfully creating a new car", func() {
		It("should return the correct response", func() {
			requestBody := `{"id": "701","title": "GM","color": "Transparent"}`
			request, _ := http.NewRequest("POST", "http://localhost:8081/cars", bytes.NewBuffer([]byte(requestBody)))
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
		})
	})

	Context("When creating an invalid car", func() {
		It("should return an error in the response", func() {
			requestBody := ""
			request, _ := http.NewRequest("POST", "http://localhost:8081/cars", bytes.NewBuffer([]byte(requestBody)))
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusBadRequest))

			responseBody, _ := io.ReadAll(response.Body)
			Expect(string(responseBody)).To(ContainSubstring("Failed"))
		})
	})
})

var _ = Describe("Car retrieval test", func() {
	// Context for testing the GET request on cars
	Context("When retrieving cars", func() {
		It("should return a list of cars", func() {
			// First create a car to ensure there is at least one car available for retrieval
			createRequestBody := `{"id": "701","title": "GM","color": "Transparent"}`
			createRequest, _ := http.NewRequest("POST", "http://localhost:8081/cars", bytes.NewBuffer([]byte(createRequestBody)))
			createRequest.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			_, err := client.Do(createRequest)
			Expect(err).NotTo(HaveOccurred())

			// Perform the GET request
			getRequest, _ := http.NewRequest("GET", "http://localhost:8081/cars", nil)
			getResponse, err := client.Do(getRequest)
			Expect(err).NotTo(HaveOccurred())
			Expect(getResponse.StatusCode).To(Equal(http.StatusOK))

			// Check the response body contains the car
			responseBody, _ := io.ReadAll(getResponse.Body)
			Expect(string(responseBody)).To(ContainSubstring("GMC"))
		})
	})
})
