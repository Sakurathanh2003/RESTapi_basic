package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var students = []student{
	{ID: "1", Name: "Thanh"},
	{ID: "2", Name: "Oanh"},
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func addStudent(c *gin.Context) {
	var newStudent student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, students)
}

func getStudentByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range students {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/students", addStudent)
	router.GET("/students/:id", getStudentByID)

	router.Run("localhost:8080")
}
