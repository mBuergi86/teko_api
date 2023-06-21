package Entity

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

type JsonExtendedResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonDataResponse struct {
	Data []Todo `json:"data"`
}

type JsonErrorResponse struct {
	Error ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}
