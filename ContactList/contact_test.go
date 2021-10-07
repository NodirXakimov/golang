package main

import "testing"

func TestGetContacts(t *testing.T) {
	//
	if len(getContacts()) != 2 {
		t.Error("Test failed........", len(getContacts()))
	}
}

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
	if len(contacts) != 3 {
		t.Error("Length of the contacts variable must be 1 after inserting first element to it")
	}

	// Checking returned and created Contact objects
	contact.ID = "3"
	if contact != returnedContact {
		t.Error("Sent and Recieved objects aren't equal")
	}

}

func TestGetContact(t *testing.T) {
	// Getting second element of contact list (which is Adham)
	contact, err := getContact(2)
	if err != nil {
		t.Error("Second element should be exist")
	}
	if contact.FirstName != "Adham" {
		t.Error("First name of the second contact must be Adham")
	}
}

func TestUpdateContact(t *testing.T) {
	// Creating new Contact object to update second element
	contact := Contact{
		FirstName: "Xusan",
		LastName:  "Xakimov",
		Phone:     "+998885550808",
		Email:     "xusanxakimov@gmail.com",
		Position:  "Main manager",
	}
	c, er := updateContact(2, contact)
	if er != nil {
		t.Error("Second element of array should be exist")
	}
	if c.ID != "2" {
		t.Error("ID of the returned contact should be similar with sent contact")
	}
	c, er = updateContact(4, contact)
	if er == nil {
		t.Error("It should return error if sent not exist element to update")
	}
}

func TestDeleteContact(t *testing.T) {
	err := deleteContact(2)
	if err != nil {
		t.Error("It should be possible delete second element of contacts slice")
	}
	err = deleteContact(4)
	if err == nil {
		t.Error("It shouldn't be possible delete which is not exist")
	}
}
