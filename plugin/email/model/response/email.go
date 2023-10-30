package response

type Email struct {
	To      string `json:"to"`      // To Email
	Subject string `json:"subject"` // Subject
	Body    string `json:"body"`    // Body
}
