package model

// Todo holds the properties of a todo in the database.
type Todo struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}
