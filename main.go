package main

import (
	"github.com/gin-gonic/gin" // Import the packages
	"net/http"                 // needed for your code.
)

type Car struct { // Declaring Car struct.
	ID    string `json:"id"`    // "json" tag specifies what a field’s
	Title string `json:"title"` // name should be when the struct’s
	Color string `json:"color"` // contents are serialized into JSON.
}

var Cars []Car // Creating an empty slice of Car struct.

func main() {
	Cars = []Car{ // Declaring a data,
		{ID: "1", Title: "BMW", Color: "Black"}, // which our Cars slice
		{ID: "2", Title: "Tesla", Color: "Red"}, // will contain of.
	}

	router := gin.Default()               // Initializing a Gin router.
	router.GET("/cars", getCars)          // Assigning handler's
	router.GET("/cars/:id", getCarByID)   // functions to
	router.POST("/cars", createCar)       // different
	router.DELETE("/cars/:id", deleteCar) // endpoint paths.

	router.Run("localhost:8081") // Attaching the router to http.Server and starting the server.
}

func createCar(c *gin.Context) {
	var newCar Car                              // Declaring a new Car variable.
	if err := c.BindJSON(&newCar); err != nil { // Bind the request body to newCar variable.
		c.IndentedJSON(http.StatusBadRequest, // Return 400 Status Code,
			gin.H{"message": "Failed to create a car"}) // and error message if binding has failed.
		return
	}

	Cars = append(Cars, newCar)                // Append the received Car struct to the Cars slice.
	c.IndentedJSON(http.StatusCreated, newCar) // Add status code with JSON representing new car to the response.
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Cars) // Add status code with JSON representing the Cars slice.
}

func getCarByID(c *gin.Context) {
	for _, car := range Cars { // Loop through Cars slice.
		if car.ID == c.Param("id") { // Extract the ID in the request path and locate a car that matches.
			c.IndentedJSON(http.StatusOK, car) // Add status code with JSON representing a car found by provided ID.
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, // Return 400 Status Code and
		gin.H{"message": "Requested car is not found"}) // error message if no car was found.
}

func deleteCar(c *gin.Context) {
	for index, car := range Cars { // Loop through Cars slice.
		if car.ID == c.Param("id") { // Extract the ID in the request path and locate a car that matches.
			Cars = append(Cars[:index], Cars[index+1:]...) // Delete found car from slice.
			c.IndentedJSON(http.StatusOK,                  // Return 400 Status Code,
				gin.H{"message": "Car is deleted"}) // and success message.
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, // Return 400 Status Code and
		gin.H{"message": "Car is not found"}) // error message if no car was found.
}
