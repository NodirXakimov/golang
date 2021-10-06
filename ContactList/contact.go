package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Contact struct (Model)
type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Position  string `json:"position"`
}

var contacts []Contact

// Get all contacts
func getContacts() []Contact {
	return contacts
}

// Get single contact
func getContact(id int) {
	// // Loop through contacts and find by id
	// for _, contact := range contacts {
	// 	if contact.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(contact)
	// 		return
	// 	}
	// }
}

// Create new contact
func (contact *Contact) createContact() (c Contact, err error) {
	c = *contact
	c.ID = strconv.Itoa(len(contacts) + 1) // Mock ID - not safe
	contacts = append(contacts, c)
	return
}

// Update a contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = params["id"] // Mock ID - not safe
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Delete a contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

func main() {

	// Mock data
	contacts = append(contacts,
		Contact{"1", "Nodir", "Xakimov", "+998999905518", "nodirxakimov@yandex.ru", "Go developer"},
		Contact{"2", "Adham", "Xakimov", "+998997401520", "adhamxakimov@gmail.com", "Project manager"},
	)

	// Getting all contacts
	fmt.Println(getContacts())

	// Creating new contact
	contact := Contact{
		FirstName: "Xasan",
		LastName:  "Xakimov",
		Phone:     "+998977775544",
		Email:     "xasanxakimov@gmail.com",
		Position:  "Investor",
	}
	fmt.Println(contact.createContact())
}
