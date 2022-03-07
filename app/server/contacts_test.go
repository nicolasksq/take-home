package server_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"server/app/dao"
	"server/app/server/responses"
)

const (
	firstName = "Nico"
	lastName  = "Andreoli"
	email     = "nicolasandreoli@gmail.com"
)

func (suite *ServerTestSuite) TestSyncContacts_happyPath() {
	//setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := suite.mockedServer()
	contactsDao := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}

	//expected
	expectedResponse := responses.SyncResponse{
		SyncedContacts: len(contactsDao),
		Contacts: []responses.Contact{{
			Firstname: firstName,
			Lastname:  lastName,
			Email:     email,
		}},
	}
	expectedCode := http.StatusOK

	suite.mockAPI.On("SyncContacts").Return(contactsDao, nil)

	s.SyncContacts(c)

	b, _ := ioutil.ReadAll(w.Body)
	var actualBody responses.SyncResponse
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	suite.Equal(expectedCode, w.Code)
	suite.Equal(expectedResponse, actualBody)
	suite.mockAPI.AssertExpectations(suite.T())
}

func (suite *ServerTestSuite) TestSyncContacts_err() {
	//setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := suite.mockedServer()

	//expected
	err := errors.New("something went wrong")
	expectedResponse := responses.SyncResponse{Error: err.Error()}
	expectedCode := http.StatusInternalServerError

	suite.mockAPI.On("SyncContacts").Return(nil, err)

	s.SyncContacts(c)

	b, _ := ioutil.ReadAll(w.Body)
	var actualBody responses.SyncResponse
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	suite.Equal(expectedCode, w.Code)
	suite.Equal(expectedResponse, actualBody)
	suite.mockAPI.AssertExpectations(suite.T())
}
