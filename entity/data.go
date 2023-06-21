package entity

import "strconv"

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Terminated  bool   `json:"terminated"`
}

func (t Todo) Serialize() []string {
	todoSerialized := []string{t.Id, t.Title, t.Description, strconv.FormatBool(t.Terminated)}
	return todoSerialized
}

type CSVData struct {
	Todo []Todo
}
