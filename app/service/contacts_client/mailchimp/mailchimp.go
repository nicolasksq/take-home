package mailchimp

import (
	"net/http"
	"server/app/service/contacts_client"
	"server/app/vendors/mailchimp-go"
)

var _ contacts_client.ClientAPI = Mailchimp{}

// Mailchimp manages communication with the Mailchimp API.
type Mailchimp struct{}

func NewMailchimp(apiKey string, httpClient *http.Client) Mailchimp {
	_ = mailchimp.SetKey(apiKey)
	mailchimp.SetClient(httpClient)
	return Mailchimp{}
}
