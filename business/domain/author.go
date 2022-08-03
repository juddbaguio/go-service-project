package domain

type Author struct {
	ID        int    `json:"authorId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
