package mailchimp

import "server/app/service/email_tool_client"

var _ email_tool_client.ClientAPI = Mailchimp{}

// Mailchimp manages communication with the Mailchimp API.
type Mailchimp struct{
	lib MailchimpLibImpl
}

func NewMailchimp(mailchimpConfig MailchimpLibImpl) Mailchimp {
	return Mailchimp{mailchimpConfig}
}