package responses

type SyncResponse struct {
	SyncedContacts int       `json:"syncedContacts"`
	Contacts       []Contact `json:"contacts"`
}

type Contact struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
