package errors

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error:   true,
		Message: message,
	}
}

var (
	//ErrInvalidID      = NewErrorResponse("Invalid ID")
	ErrTodoNotFound   = NewErrorResponse("Todo not found")
	ErrFileNotCreated = NewErrorResponse("File cannot be created")
	ErrNotImplemented = NewErrorResponse("Not implemented")
)
