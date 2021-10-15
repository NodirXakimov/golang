package main

import (
	"fmt"
	"golang/ContactListDB/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Contact struct {
	ID        string  `json:"id" db:"id"`
	FirstName string  `json:"firstname" db:"firstname"`
	LastName  string  `json:"lastname" db:"lastname"`
	Phone     string  `json:"phone" db:"phone"`
	Email     *string `json:"email" db:"email"`
	Position  string  `json:"position" db:"position"`
}

// Getting all contacts form DB
func getContacts(c *gin.Context) {
	db := config.Connect()
	contacts := []Contact{}
	err := db.Select(&contacts, "SELECT * FROM contacts;")
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, err)
		fmt.Println("error: ", err)
	}
	c.IndentedJSON(http.StatusOK, contacts)
}

// Getting single contact from DB
func getContact(c *gin.Context) {
	db := config.Connect()
	contact := Contact{}
	id := c.Param("id")
	err := db.Get(&contact, "SELECT * FROM contacts WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
	} else {
		c.IndentedJSON(http.StatusOK, contact)
	}
}

// Storing new element to the DB
func createContact(c *gin.Context) {
	db := config.Connect()

	contact := Contact{}

	// Call BindJSON to bind the received JSON to task.
	if err := c.BindJSON(&contact); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		if contact.FirstName == "" || contact.LastName == "" || contact.Phone == "" || contact.Position == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Credentials not valid", "credentials": "firstname, lastname, email*, phone, position"})
			return
		}
		newUserQuery := "INSERT INTO contacts (firstname, lastname, phone, email, position) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
		err := db.QueryRow(newUserQuery, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position).Scan(&contact.ID)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, contact)
		} else {
			c.IndentedJSON(http.StatusCreated, err)
		}
	}

}

func main() {

	router := gin.Default()
	router.GET("/api/contacts", getContacts)
	router.POST("/api/contacts", createContact)
	router.GET("/api/contacts/:id", getContact)
	// router.PUT("/api/tasks/:id", updateTask)
	// router.DELETE("/api/tasks/:id", deleteTask)

	log.Fatal(router.Run(":1212"))
}
