package mailchimp_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"server/app/dao"
	"server/app/service/email_tool_client/mailchimp"
	"server/app/service/email_tool_client/mailchimp/mocks"
	"server/app/vendors/mailchimp-go/lists"
	"server/app/vendors/mailchimp-go/lists/members"
)

const (
	firstName = "nico"
	lastName  = "andreoli"
	email     = "nicolasandreoli9@gmail.com"
)

type MailchimpTestSuite struct {
	suite.Suite
	m *mocks.MailchimpLib
}

func TestMailchimpTestSuite(t *testing.T) {
	suite.Run(t, new(MailchimpTestSuite))
}

func (ms *MailchimpTestSuite) SetupTest() {
	ms.m = new(mocks.MailchimpLib)
}

func (ms *MailchimpTestSuite) TestBatchListMembers_happyPath() {
	mailchimp := mailchimp.NewMailchimp(ms.m)
	membersResponse := &members.MembersResponse{
		NewMembers: []members.MemberResponse{{EmailAddress: email}},
		Errors:     nil,
		ErrorCount: 0,
	}

	params := &members.NewParams{
		Members: []members.Member{{
			EmailAddress: email,
			Status:       members.StatusSubscribed,
		}},
		UpdateExisting: true,
	}

	contacts := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}

	ms.m.On("NewMembers", "nicolas.andreoli", params).Return(membersResponse, nil)
	result, err := mailchimp.BatchListMembers(contacts, "nicolas.andreoli")

	ms.m.AssertExpectations(ms.T())
	ms.Equal(contacts, result)
	ms.Nil(err)
}

func (ms *MailchimpTestSuite) TestBatchListMembers_error() {
	mailchimp := mailchimp.NewMailchimp(ms.m)
	membersResponse := &members.MembersResponse{
		NewMembers: []members.MemberResponse{{EmailAddress: email}},
		Errors: []members.Error{{
			EmailAddress: email,
			Error:        "some weird error",
		}},
		ErrorCount: 1,
	}

	params := &members.NewParams{
		Members: []members.Member{{
			EmailAddress: email,
			Status:       members.StatusSubscribed,
		}},
		UpdateExisting: true,
	}

	contacts := []dao.Contact{{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}

	var expectedResult = []dao.Contact{}
	ms.m.On("NewMembers", "nicolas.andreoli", params).Return(membersResponse, nil)
	result, err := mailchimp.BatchListMembers(contacts, "nicolas.andreoli")

	ms.m.AssertExpectations(ms.T())
	ms.Equal(expectedResult, result)
	ms.Nil(err)
}

func (ms *MailchimpTestSuite) TestGetListsByName_happyPath() {
	mailchimp := mailchimp.NewMailchimp(ms.m)

	getParams := &lists.GetParams{
		Fields: []string{"lists.id", "lists.name"},
	}
	listsResponse := &lists.Lists{
		Lists: []lists.List{{
			ID:   "id_list",
			Name: "nicolas.andreoli",
		}},
		TotalItems: 1,
	}

	listExpected := &dao.List{
		ID:   "id_list",
		Name: "nicolas.andreoli",
	}

	ms.m.On("GetLists",  getParams).Return(listsResponse, nil)
	result, err := mailchimp.GetListsByName("nicolas.andreoli")

	ms.m.AssertExpectations(ms.T())
	ms.Equal(listExpected, result)
	ms.Nil(err)
}
