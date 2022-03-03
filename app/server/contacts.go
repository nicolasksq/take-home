package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"server/app/api/contacts"
	"server/app/server/request"
	"server/app/server/responses"
)

func SyncContacts(c *gin.Context) {
	ct := contacts.NewContacts(os.Getenv(apiKeyClient))
	res, err := ct.SyncContacts()
	if err != nil {
		switch err.(type) {
			// error not found case
			// error bad request
			// error ...
		}
		// we could add a different kind of errors in the future depending of err type
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	response := responses.SyncResponse{
		SyncedContacts: len(res),
	}

	contactsList := make([]responses.Contact, len(res))
	for i := range res {
		contactsList[i].Firstname = res[i].Firstname
		contactsList[i].Lastname = res[i].Lastname
		contactsList[i].Email = res[i].Email
	}
	response.Contacts = contactsList

	c.IndentedJSON(http.StatusOK, response)
}

// CreateList dummy endpoint to create the list for first time, it's prepared to receive same format as mailchimp.
// but we are using default values as requested.
func CreateList(c *gin.Context) {
	ct := contacts.NewContacts(os.Getenv(apiKeyClient))

	var requestBody *request.CreateList
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "missing list name"})
		return
	}

	err := ct.CreateList(&requestBody.ListName)
	if err != nil {
		// we could add a different kind of errors in the future depending of err type
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
