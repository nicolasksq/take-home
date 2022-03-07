# Back-End Project Trio

****

This project is version 0.0.1, and it was developed in Go 1.16.

The take-home project has 1 main endpoint which is in charge of sync contacts between MockAPI and Mailchimp.

# Table Content

- [Structure](#Structure)
- [Description](#Description)
- [Resources](#Resources)
- [Endpoints](#Endpoints)
- [Technical Design](#Technical Design)
- [Access Token (API Key)](#Access)


# Structure project

### A top-level directory layout 
it doesn't have all files and folders, just the important ones to understand how the project was build 

    .
    ├── app    # application folder
    ├────── api  # bussiness logic application
    ├────── dao  # data access object folder
    ├────── server    
    ├──────── ....
    ├────── service   # it has third-party clients.
    ├──────── contacts_client   # I called contacts client for seems to be an external client. It will handle each client that we want to add here.
    ├────────── mailchimp         # mailchimp folder
    ├───────────── list.go          # it has the list methods for mailchimp
    ├───────────── mailchimp.go     # it has the instance for mailchimp
    ├────────── client.go     # Interface that must be implemented in contacts folder
    ├──────── contacts_service   # I called contacts service cuz seems to be a local/internal service.
    ├────── vendors   # this folder has been created just to lift one lib to handle mailchimp, and it was modified to satisfy the necessities. It shoudn't be there, we should have our own lib for mailchimp, or keep searching other better skd for Go. 
    ├── cmd           # Main applications for this project.
    ├── go.mod    
    └── README.md

# Description project

The project syncs a contact list from mockAPI to Mailchimp. In order to push the contacts, we've already created the list needed.

If we want to test the behaviour running the endpoint without a list, we should either delete it using a mailchimp endpoint or changing the [DefaultListName][https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L10]

If we want to create a new list using the endpoint provided, you need to be sure that there is no list created, because our account just allows to create one list.

After having a list created, you'll be able to sync the contacts, using the endpoint provided.

There was a couple of unknowns here, such as
- should we update the contact if is already create? you will find this logic [here][https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L91]
- [we are searching the list ID calling `GET /lists` endpoint][https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L43]. We expect to have just one list, as it's a limitation for Mailchimp free account.


# Resources

- https://613b9035110e000017a456b1.mockapi.io/api/v1/
- https://`{{$PREFIX_TOKEN}}`.api.mailchimp.com/3.0/

HOST: `https://afternoon-citadel-79267.herokuapp.com/`

#  Endpoints
- POST `/list`
  - Request: 
  ```json 
    {
        "listName":"nicolas.andreoli"
    }
    ```
  - Response:
    `Status code: 201`

- GET `/`
    - response: fun stuff
    

- GET `/contacts/sync`
    - Response:
```json 
{
    "syncedContacts": 1,
    "contacts": [{
            "firstname": "Michelle",
            "lastname": "Gaylord",
            "email": "Kirk.Fritsch93@hotmail.com"
        }]
}
```
  - Error response:
    
      This error is caused when we try to push contacts to Mailchimp and the list has not been created.
    
```json 
{
	"syncedContacts": 0,
	"contacts": [],
	"error": "something went wrong or there is no list with the given name"
}
```
    
***

# Technical Design

The main goal is to sync two different API's, integrating contacts from one MockAPI and pushing them to Mailchimp.
In order to do that, I decided to create a `service/contacts_service` folder and `service/email_tool_client`, which will have an interface.
```go 
type ContactAPI interface {
    GetContacts() ([]dao.Contact, error)
}

type ClientAPI interface {
    BatchListMembers(contacts []dao.Contact, listID string) ([]dao.Contact, error)
    GetListsByName(name string) (*dao.List, error)
    CreateList(listName *string) error
}
```

These interfaces need to be implemented by the service we decide to use. So, this give us the flexibility to implement another API instead of MockAPI or Mailchimp without changing the main business logic.

To have the communication to either Mailchimp or other email-tool, we need to set the `apikey` as an env variable.

-----

We have two DataAccessObjects, such as `contact, list`,  to isolate the application model from the others API's.
```go 
type Contact struct {
  FirstName string
  LastName  string
  Email     string
}

type List struct {
	ID   string
	Name string
}
```

The `mailchimp_lib.go` in `service/email_tool_client/mailchimp` has been created to mock the library that I downloaded. 
I use this approach to make static methods testable in our business.

# Access Tokens

The Token to play with Mailchimp API: `d2737a48f10fa2e974bc1ab2a2cd22bd-us14`

[https://github.com/nicolasksq/take-home-trio/blob/master/app/service/email_tool_client/mailchimp/list.go#L10]: https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L10

[here]: https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L91

[https://github.com/nicolasksq/take-home-trio/blob/master/app/service/email_tool_client/mailchimp/list.go#L91]: https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L91

[https://github.com/nicolasksq/take-home-trio/blob/master/app/service/email_tool_client/mailchimp/list.go#L43]: https://github.com/nicolasksq/take-home-trio/blob/master/app/service/contacts_client/mailchimp/list.go#L43
