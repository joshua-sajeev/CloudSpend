package errors

type ServiceError struct {
	Code    int
	Message string
	Details string
}

func NewServiceError(code int, message string) *ServiceError
func (e *ServiceError) Error() string
func HandleServiceError(err error) *ServiceError
