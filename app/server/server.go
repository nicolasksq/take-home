package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

const HOST = "localhost"
const PORT = ":80"

const defaultEndpoint = "/"
const contactsSyncEndpoint = "/contacts/sync"
const listEndpoint = "/list/"

type Server struct{}

const apiKeyClient = "apikey"

func (s Server) Start() error {
	// there is a better way to do this, using flag package.
	for i := range os.Args {
		if strings.Contains(os.Args[i], "--"+apiKeyClient+"=") {
			_ = os.Setenv(apiKeyClient, strings.Trim(os.Args[i], "--"+apiKeyClient+"="))
		}
	}

	if os.Getenv(apiKeyClient) == "" {
		panic("apikey must to be set")
	}

	router := gin.Default()
	// we defined endpoints here
	router.GET(defaultEndpoint, func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"trio.dev": "Hey Trio team!, hope you're having a nice day. :)"})
	})

	// main endpoint to sync contacts from the given service to given api
	router.GET(contactsSyncEndpoint, SyncContacts)

	// main endpoint to sync contacts from the given service to given api
	router.POST(listEndpoint, CreateList)

	if err := router.Run(HOST + PORT); err != nil {
		return err
	}
	return nil
}
