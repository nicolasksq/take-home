package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/app/server/request"
	"server/app/server/responses"
)

func (s Server) SyncContacts(c *gin.Context) {
	res, err := s.ContactsAPI.SyncContacts()
	if err != nil {
		switch err.(type) {
		// error not found case
		// error bad request
		// error ...
		}
		// we could add a different kind of errors in the future depending of err type
		c.IndentedJSON(http.StatusInternalServerError, responses.SyncResponse{Error: err.Error()})
		return
	}

	response := responses.SyncResponse{
		SyncedContacts: len(res),
	}

	contactsList := make([]responses.Contact, len(res))
	for i := range res {
		contactsList[i].Firstname = res[i].FirstName
		contactsList[i].Lastname = res[i].LastName
		contactsList[i].Email = res[i].Email
	}
	response.Contacts = contactsList

	c.IndentedJSON(http.StatusOK, response)
}

// CreateList dummy endpoint to create the list for first time, it's prepared to receive same format as mailchimp.
// but we are using default values as requested.
func (s Server) CreateList(c *gin.Context) {
	var requestBody *request.CreateList
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "missing list name"})
		return
	}

	if err := s.ContactsAPI.CreateList(&requestBody.ListName); err != nil {
		// we could add a different kind of errors in the future depending of err type
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
