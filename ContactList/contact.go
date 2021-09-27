package main

import (
	"encoding/json"
	"log"
	"math/rand"
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
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Get single contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through contacts and find by id
	for _, contact := range contacts {
		if contact.ID == params["id"] {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Create new contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
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

	// Init Router
	r := mux.NewRouter()

	// Mock data
	contacts = append(contacts,
		Contact{"1", "Nodir", "Xakimov", "+998999905518", "nodirxakimov@yandex.ru", "Go developer"},
		Contact{"2", "Adham", "Xakimov", "+998997401520", "adhamxakimov@gmail.com", "Project manager"},
	)

	// Route Handlers / Endpoints
	r.HandleFunc("/api/contacts", getContacts).Methods("GET")
	r.HandleFunc("/api/contacts/{id}", getContact).Methods("GET")
	r.HandleFunc("/api/contacts", createContact).Methods("POST")
	r.HandleFunc("/api/contacts/{id}", updateContact).Methods("PUT")
	r.HandleFunc("/api/contacts/{id}", deleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8008", r))
}
