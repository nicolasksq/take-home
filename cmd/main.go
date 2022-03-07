package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"server/app/api/contacts"
	"server/app/server"
	"server/app/service/contacts_service/mockapi"
	"server/app/service/email_tool_client/mailchimp"
)

const apiKeyClient = "apikey"
const timeout = 5 * time.Second

func main() {
	var apikey string
	if os.Getenv("apikey") != "" {
		//apikey set in heroku
		apikey = os.Getenv("apikey")
	} else {
		//local enviroment.
		for i := range os.Args {
			if strings.Contains(os.Args[i], "--"+apiKeyClient+"=") {
				apikey = strings.Trim(os.Args[i], "--"+apiKeyClient+"=")
			}
		}
	}

	log.Logger.Println(apikey)

	// we need to fail if apikey is not set.
	if apikey == "" {
		panic("apikey must to be set")
	}

	service := mockapi.NewMockapi(&http.Client{Timeout: timeout})
	emailTool := mailchimp.NewMailchimp(mailchimp.NewMailchimpLib(apikey, &http.Client{Timeout: timeout}))
	contactsAPI := contacts.NewContacts(service, emailTool)

	s := server.NewServer(contactsAPI)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
