package model

type Todo struct {
	ID    int64  `json:"id" db:"id, primarykey, autoincrement"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body"`
}
