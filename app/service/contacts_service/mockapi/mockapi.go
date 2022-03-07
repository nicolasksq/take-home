package mockapi

import (
	"encoding/json"
	"io/ioutil"

	"server/app/dao"
	"server/app/service/contacts_service"
	clientHttp "server/app/service/contacts_service/mockapi/client"
)

const mockApiEndpoint = "https://613b9035110e000017a456b1.mockapi.io/api/v1/"
const getContactsEndpoint = mockApiEndpoint + "/contacts"

var _ contacts_service.ContactAPI = Mockapi{}

type Mockapi struct {
	httpClient clientHttp.HTTPClient
}

func NewMockapi(client clientHttp.HTTPClient) Mockapi {
	// I set a timeout to call this api, in case it could be unresponsive.
	return Mockapi{httpClient: client}
}

func (m Mockapi) GetContacts() ([]dao.Contact, error) {
	res, err := m.httpClient.Get(getContactsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	var contacts []dao.Contact
	if err := json.Unmarshal(body, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}
