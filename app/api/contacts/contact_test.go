package contacts_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/suite"

	"server/app/api/contacts"
	"server/app/dao"
	serviceMock "server/app/service/contacts_service/mocks"
	clientMock "server/app/service/email_tool_client/mocks"
)

const (
	firstName = "Nico"
	lastName  = "Andreoli"
	email     = "nicolasandreoli9@gmail.com"
)

var genericError = errors.New("something went wrong")

type ContactAPITestSuite struct {
	suite.Suite
	mockClient  *clientMock.ClientAPI
	mockService *serviceMock.ContactAPI
}

func (cAPI *ContactAPITestSuite) SetupTest() {
	cAPI.mockClient = new(clientMock.ClientAPI)
	cAPI.mockService = new(serviceMock.ContactAPI)
}

func TestContactAPITestSuite(t *testing.T) {
	suite.Run(t, &ContactAPITestSuite{})
}

func (cAPI ContactAPITestSuite) TestSyncContacts_happyPath() {
	// setup
	contactsDao := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}
	cAPI.mockService.On("GetContacts").Return(contactsDao, nil).Once()

	list := &dao.List{
		ID:   "id_list",
		Name: "nicolas.andreoli",
	}
	cAPI.mockClient.On("GetListsByName", "nicolas.andreoli").Return(list, nil).Once()
	cAPI.mockClient.On("BatchListMembers", contactsDao, list.ID).Return(contactsDao, nil).Once()

	contactsAPI := contacts.NewContacts(cAPI.mockService, cAPI.mockClient)
	actualContacts, err := contactsAPI.SyncContacts()

	cAPI.mockClient.AssertExpectations(cAPI.T())
	cAPI.Nil(err)
	cAPI.Equal(contactsDao, actualContacts)
}

func (cAPI ContactAPITestSuite) TestSyncContacts_whenGetContactsFails() {
	// setup
	//expected error
	cAPI.mockService.On("GetContacts").Return(nil, genericError).Once()

	cAPI.mockClient.On("GetListsByName", mock.Anything).Return(nil, nil).Maybe()
	cAPI.mockClient.On("BatchListMembers", mock.Anything, mock.Anything).Return(nil, nil).Maybe()

	contactsAPI := contacts.NewContacts(cAPI.mockService, cAPI.mockClient)
	actualContacts, err := contactsAPI.SyncContacts()

	cAPI.mockClient.AssertExpectations(cAPI.T())
	cAPI.Nil(actualContacts)
	cAPI.NotNil(err)
	cAPI.Equal(genericError.Error(), err.Error())
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "GetListsByName", 0)
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "BatchListMembers", 0)
}

func (cAPI ContactAPITestSuite) TestSyncContacts_whenGetListsByNameFails() {
	// setup
	contactsDao := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}
	cAPI.mockService.On("GetContacts").Return(contactsDao, nil).Once()

	cAPI.mockClient.On("GetListsByName", "nicolas.andreoli").Return(nil, genericError).Once()
	cAPI.mockClient.On("BatchListMembers", mock.Anything, mock.Anything).Return(nil, nil).Maybe()

	contactsAPI := contacts.NewContacts(cAPI.mockService, cAPI.mockClient)
	actualContacts, err := contactsAPI.SyncContacts()

	cAPI.mockClient.AssertExpectations(cAPI.T())
	cAPI.Nil(actualContacts)
	cAPI.NotNil(err)
	cAPI.Equal(genericError.Error(), err.Error())
	cAPI.mockService.AssertNumberOfCalls(cAPI.T(), "GetContacts", 1)
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "GetListsByName", 1)
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "BatchListMembers", 0)
}

func (cAPI ContactAPITestSuite) TestSyncContacts_whenBatchListMembersFails() {
	// setup
	contactsDao := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}
	cAPI.mockService.On("GetContacts").Return(contactsDao, nil).Once()
	list := &dao.List{
		ID:   "id_list",
		Name: "nicolas.andreoli",
	}
	cAPI.mockClient.On("GetListsByName", "nicolas.andreoli").Return(list, nil).Once()
	cAPI.mockClient.On("BatchListMembers", mock.Anything, mock.Anything).Return(nil, genericError).Once()

	contactsAPI := contacts.NewContacts(cAPI.mockService, cAPI.mockClient)
	actualContacts, err := contactsAPI.SyncContacts()

	cAPI.mockClient.AssertExpectations(cAPI.T())
	cAPI.Nil(actualContacts)
	cAPI.NotNil(err)
	cAPI.Equal(genericError.Error(), err.Error())
	cAPI.mockService.AssertNumberOfCalls(cAPI.T(), "GetContacts", 1)
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "GetListsByName", 1)
	cAPI.mockClient.AssertNumberOfCalls(cAPI.T(), "BatchListMembers", 1)
}