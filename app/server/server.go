package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"server/app/api/contacts"
)

const HOST = "localhost"
const PORT = ":8080"

const defaultEndpoint = "/"
const contactsSyncEndpoint = "/contacts/sync"
const listEndpoint = "/list/"

type Server struct {
	ContactsAPI contacts.ContactInterface
}

func NewServer(contactsApi contacts.ContactInterface) Server {
	return Server{
		ContactsAPI: contactsApi,
	}
}

func (s Server) Start() error {
	router := gin.Default()
	// we defined endpoints here
	router.GET(defaultEndpoint, func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"hey": "hope you're having a nice day. :)"})
	})

	// main endpoint to sync contacts from the given service to given api
	router.GET(contactsSyncEndpoint, s.SyncContacts)

	// main endpoint to sync contacts from the given service to given api
	router.POST(listEndpoint, s.CreateList)

	// to satisfy heroku port
	port := os.Getenv("PORT")
	var addr string
	if port == "" {
		addr = HOST + PORT
	} else {
		addr = ":" + port
	}

	if err := router.Run(addr); err != nil {
		return err
	}
	return nil
}
