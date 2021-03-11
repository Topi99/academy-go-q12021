package controller

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Key     string `json:"error"`
	Message string `json:"message"`
}

// NewHTTPError creates nuew HTTPError
func NewHTTPError(code int, key string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Key:     key,
		Message: msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	return e.Key + ": " + e.Message
}
