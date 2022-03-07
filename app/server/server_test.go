package server_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"server/app/api/contacts/mocks"
	"server/app/server"
)

type ServerTestSuite struct {
	suite.Suite
	mockAPI *mocks.ContactInterface
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (suite *ServerTestSuite) mockedServer() server.Server {
	return server.NewServer(suite.mockAPI)
}

func (suite *ServerTestSuite) SetupTest() {
	os.Setenv("apikey", "test")
	gin.SetMode(gin.ReleaseMode)
	suite.mockAPI = new(mocks.ContactInterface)
}
