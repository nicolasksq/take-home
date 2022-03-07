package email_tool_client

import (
	"server/app/dao"
)

// ClientAPI SyncContacts() is to send given contacts to the current client
type ClientAPI interface {
	// BatchListMembers given a list of contacts, we push members to the list
	BatchListMembers(contacts []dao.Contact, listID string) ([]dao.Contact, error)
	// GetListsByName give us a list of the current list created. We make this call to get the list ID
	GetListsByName(name string) (*dao.List, error)
	CreateList(listName *string) error
}
