package contacts_service

import (
	"server/app/dao"
)

type ContactAPI interface {
	GetContacts() ([]dao.Contact, error)
}
