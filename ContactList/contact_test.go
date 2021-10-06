package main

import "testing"

// func TestGetContacts(t *testing.T) {
// 	//
// 	if len(getContacts()) != 2 {
// 		t.Error("Test failed........", len(getContacts()))
// 	}
// }

func TestCreateContact(t *testing.T) {
	// Creating test contact
	contact := Contact{
		FirstName: "Xasan",
		LastName:  "Xakimov",
		Phone:     "+998977775544",
		Email:     "xasanxakimov@gmail.com",
		Position:  "Investor",
	}

	returnedContact, _ := contact.createContact()

	// Checking length of the contacts
	if len(contacts) != 1 {
		t.Error("Length of the contacts variable must be 1 after inserting first element to it")
	}

	// Checking returned and created Contact objects
	contact.ID = "1"
	if contact != returnedContact {
		t.Error("Sent and Recieved objects aren't equal")
	}

}
