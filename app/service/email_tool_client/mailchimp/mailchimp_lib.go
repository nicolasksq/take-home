package mailchimp

import (
	"net/http"

	"server/app/vendors/mailchimp-go/lists"
	"server/app/vendors/mailchimp-go/lists/members"
)

// NewMailchimpLib we create this to make testeable the mailchimlib in vendors folder.
func NewMailchimpLib(apiKey string, httpClient *http.Client) MailchimpLibImpl {
	return MailchimpLib{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

type MailchimpLib struct {
	apiKey     string
	httpClient *http.Client
}

// MailchimpLibImpl due to Kata test
type MailchimpLibImpl interface {
	GetLists(params *lists.GetParams) (*lists.Lists, error)
	NewMembers(listID string, params *members.NewParams) (*members.MembersResponse, error)
	NewList(params *lists.NewParams) (*lists.List, error)
}

func (mc MailchimpLib) GetLists(params *lists.GetParams) (*lists.Lists, error) {
	return lists.Get(params)
}

func (mc MailchimpLib) NewMembers(listID string, params *members.NewParams) (*members.MembersResponse, error) {
	return members.New(listID, params)
}

func (mc MailchimpLib) NewList(params *lists.NewParams) (*lists.List, error) {
	return lists.New(params)
}
