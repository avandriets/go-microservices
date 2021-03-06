package controller

import (
	"../messsages"
	"../model"
	"log"
)

func AddContact(c *messages.Contact) (*messages.Contact, error) {

	log.Printf("INIT %d %s %s %s", c.Id, c.Name, c.Email, c.Phone)
	if contact := getContactById(c.Id); contact == nil {
		var contactId int32
		err := model.GetDatabase().QueryRow(`INSERT INTO csvdata(id, name, email, phone) VALUES($1, $2, $3, $4) RETURNING id`,
			c.Id, c.Name, c.Email, c.Phone).Scan(&contactId)
		if err == nil {
			newContact := getContactById(contactId)
			log.Printf("GET %d %s %s %s", newContact.Id, newContact.Name, newContact.Email, newContact.Phone)
			return newContact, nil
		} else {
			return nil, err
		}
	} else {
		return contact, nil
	}
}

func getContactById(id int32) *messages.Contact {
	log.Printf("GET %d", id)
	contact := messages.Contact{}
	err := model.GetDatabase().QueryRow("SELECT id, name, email, phone  FROM public.csvdata WHERE id = $1", id).Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &contact
}
