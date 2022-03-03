package mockapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/app/dao"
	"time"
)

const mockApiEndpoint = "https://613b9035110e000017a456b1.mockapi.io/api/v1/"
const getContactsEndpoint = mockApiEndpoint + "/contacts"
const timeout = 5 * time.Second

type Mockapi struct {
	httpClient http.Client
}

func NewMockapi() Mockapi {
	// I set a timeout to call this api, in case it could be unresponsive.
	return Mockapi{httpClient: http.Client{Timeout: timeout}}
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
