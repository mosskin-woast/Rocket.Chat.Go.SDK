package models

// VisitorInfo is data to register a guest
type VisitorInfo struct {
	Token      string `json:"token"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

// Visitor is a chat visitor
type Visitor struct {
	UserID        string         `json:"userId"`
	Name          string         `json:"name"`
	Token         string         `json:"token"`
	Username      string         `json:"username"`
	VisitorEmails []VisitorEmail `json:"visitorEmails"`
}

// VisitorEmail is an email address for a visitor
type VisitorEmail struct {
	Address string `json:"address"`
}
