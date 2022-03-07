package responses

type SyncResponse struct {
	SyncedContacts int       `json:"syncedContacts"`
	Contacts       []Contact `json:"contacts,omitempty"`
	Error          string    `json:"error,omitempty"`
}

type Contact struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
}
