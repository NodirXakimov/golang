package main

import (
	"errors"
	"fmt"
	"strconv"
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

var contacts = []Contact{
	{"1", "Nodir", "Xakimov", "+998999905518", "nodirxakimov@yandex.ru", "Go developer"},
	{"2", "Adham", "Xakimov", "+998997401520", "adhamxakimov@gmail.com", "Project manager"},
}

// Get all contacts
func getContacts() []Contact {
	return contacts
}

// Get single contact
func getContact(id int) (c Contact, err error) {
	// // Loop through contacts and find by id
	for _, c := range contacts {
		if c.ID == strconv.Itoa(id) {
			return c, nil
		}
	}
	return Contact{}, errors.New("Contact not found")
}

// Create new contact
func (contact *Contact) createContact() (c Contact, err error) {
	c = *contact
	c.ID = strconv.Itoa(len(contacts) + 1) // Mock ID - not safe
	contacts = append(contacts, c)
	return
}

// Update a contact
func updateContact(id int, contact Contact) (Contact, error) {
	for index, c := range contacts {
		if c.ID == strconv.Itoa(id) {
			contacts = append(contacts[:index], contacts[index+1:]...)
			contact.ID = strconv.Itoa(id)
			contacts = append(contacts, contact)
			return contact, nil
		}
	}
	return Contact{}, errors.New("Contact not found with this id")
}

// Delete a contact
func deleteContact(id int) error {
	for index, contact := range contacts {
		if contact.ID == strconv.Itoa(id) {
			contacts = append(contacts[:index], contacts[index+1:]...)
			return nil
		}
	}
	return errors.New("Contact not found")
}

func main() {

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
	fmt.Println(getContact(3))
	c, er := updateContact(3, Contact{
		FirstName: "Xusan",
		LastName:  "Xakimov",
		Phone:     "+998977775544",
		Email:     "xusanxakimov@gmail.com",
		Position:  "Investor",
	})
	fmt.Println(c, er)
	deleteContact(2)
	fmt.Println(getContacts())
}
