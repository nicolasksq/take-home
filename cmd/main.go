package main

import (
	"fmt"
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
	for i := range os.Args {
		fmt.Printf("este es el args : %+v",os.Args[i])

		if strings.Contains(os.Args[i], "--"+apiKeyClient+"=") {
			apikey = strings.Trim(os.Args[i], "--"+apiKeyClient+"=")
		}
	}


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
