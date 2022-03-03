package contacts

import (
	"net/http"
	"server/app/dao"
	"server/app/service/contacts_client"
	"server/app/service/contacts_client/mailchimp"
	"server/app/service/contacts_service"
	"server/app/service/contacts_service/mockapi"
	"time"
)

type Contacts struct {
	// client is an instance of any kind of client that we want to add.
	// this give us the flexibility to add, or change to another client instead of mailchimp
	client contacts_client.ClientAPI
	// service is an instance of any kind of service related with contacts, today is MockAPI, tomorrow could be a real service to get contats
	// which will need to have same behavior, have to use same interface
	service contacts_service.ContactAPI
}

func NewContacts(clientApiKey string) Contacts {
	// here we can decide which dependencies we want to use. So, if we want to use another client, or service
	// we must to create similar file implementing the given interface.
	return Contacts{
		service: mockapi.NewMockapi(),
		client:  mailchimp.NewMailchimp(clientApiKey, &http.Client{Timeout: 5 * time.Second}),
	}
}

func (c Contacts) SyncContacts() ([]dao.Contact, error) {
	// get contacts from mockAPI
	contactList, err := c.service.GetContacts()
	if err != nil {
		return nil, err
	}

	// we have just one list name
	l, err := c.client.GetListsByName(mailchimp.DefaultListID)
	if err != nil {
		return nil, err
	}

	// save a  []members into a list
	list, err := c.client.BatchListMembers(contactList, l.ID)
	if err != nil {
		return nil, err
	}

	// return contacts if not fail
	return list, nil
}

// CreateList just dummy logic to create a default list.
func (c Contacts) CreateList(listName *string) error {
	// create a list to start saving members
	return c.client.CreateList(listName)
}
