package mockapi_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"server/app/dao"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"server/app/service/contacts_service/mockapi"
	"server/app/service/contacts_service/mockapi/client/mocks"
)

type MockApiTestSuite struct {
	suite.Suite
	httpClient *mocks.HTTPClient
}

func TestMockApiTestSuite(t *testing.T) {
	suite.Run(t, new(MockApiTestSuite))
}

func (ms *MockApiTestSuite) SetupTest() {
	ms.httpClient = new(mocks.HTTPClient)
}

func (ms *MockApiTestSuite) TestGetContacts_happyPath() {
	json := `[{"firstName": "Nicolas","lastName": "Andreoli","email": "nicolasandreoli9@gmail.com","avatar": "https://cdn.fakercloud.com/avatars/dshster_128.jpg","id": "115"}]`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	response := &http.Response{
		Status:     "OK",
		StatusCode: http.StatusOK,
		Body:       r,
	}
	ms.httpClient.On("Get", mock.Anything).Return(response, nil)

	expectedContacs := []dao.Contact{{
		FirstName: "Nicolas",
		LastName:  "Andreoli",
		Email:     "nicolasandreoli9@gmail.com",
	}}

	mapi := mockapi.NewMockapi(ms.httpClient)
	actualResponse, err := mapi.GetContacts()
	ms.Nil(err)
	ms.Equal(expectedContacs, actualResponse)
}

func (ms *MockApiTestSuite) TestGetContacts_err() {
	e := errors.New("something went wrong")
	ms.httpClient.On("Get", mock.Anything).Return(nil, e).Once()
	mapi := mockapi.NewMockapi(ms.httpClient)

	actualResponse, err := mapi.GetContacts()
	ms.Nil(actualResponse)
	ms.NotNil(err)
	ms.Equal(e.Error(), err.Error())
}
