package mailchimp

import (
	"errors"

	"server/app/dao"
	"server/app/vendors/mailchimp-go/lists"
	"server/app/vendors/mailchimp-go/lists/members"
)

const DefaultListID = "nicolas.andreoli"

func (mc Mailchimp) CreateList(listName *string) error {
	if listName == nil {
		dl := DefaultListID
		listName = &dl
	}

	params := &lists.NewParams{
		Name:               *listName,
		Visibility:         lists.VisibilityPublic,
		PermissionReminder: "You opted to receive updates",
		Contact: &lists.Contact{
			Company:  "Trio",
			Address1: "123 Main St",
			City:     "Chicago",
			State:    "IL",
			Zip:      "60613",
			Country:  "United States",
		},
		CampaignDefaults: &lists.CampaignDefaults{
			FromName:  "Nicolas",
			FromEmail: "nicolasandreoli9@gmail.com",
			Subject:   "Trio project",
			Language:  "EN",
		},
		EmailTypeOption: false,
	}
	_, err := mc.lib.NewList(params)
	return err
}

// GetListsByName we return the first list by name. We expect to have just 1 to 1
func (mc Mailchimp) GetListsByName(name string) (*dao.List, error) {
	//since we are going to have just one list, there is no necessity to handle the paginator.
	getParams := lists.GetParams{
		Fields: []string{"lists.id", "lists.name"},
	}

	response, err := mc.lib.GetLists(&getParams)
	if err != nil {
		return nil, err
	}

	for _, list := range response.Lists {
		if list.Name == name {
			return &dao.List{
				ID:   list.ID,
				Name: list.Name,
			}, nil
		}
	}

	return nil, errors.New("something went wrong or there is no list with the given name")
}

// BatchListMembers given a list of contact, we push all of them to the given list
func (mc Mailchimp) BatchListMembers(contacts []dao.Contact, listID string) ([]dao.Contact, error) {
	m := make([]members.Member, len(contacts))
	contactMap := make(map[string]dao.Contact, len(contacts))
	for i := range contacts {
		contactMap[contacts[i].Email] = contacts[i]
		m[i].EmailAddress = contacts[i].Email
		m[i].Status = members.StatusSubscribed
	}

	params := &members.NewParams{
		Members: m,
		// TODO we don't have an specification if we have to make an update.
		// So, we're going to make it to have always a number of synced contacts.
		UpdateExisting: true,
	}

	// Add members to a list
	membersResponse, err := mc.lib.NewMembers(listID, params)
	if err != nil {
		return nil, err
	}

	// we check if there is errors
	if membersResponse.ErrorCount > 0 {
		for i := range membersResponse.Errors {
			if _, ok := contactMap[membersResponse.Errors[i].EmailAddress]; ok {
				delete(contactMap, membersResponse.Errors[i].EmailAddress)
			}
		}

		// we recreate the slice to have a new one without members with errors.
		contacts = []dao.Contact{}
		for i := range contactMap {
			contacts = append(contacts, contactMap[i])
		}
	}
	return contacts, nil
}
