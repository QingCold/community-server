package model

type Complaint struct {
	ID      int64  `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
