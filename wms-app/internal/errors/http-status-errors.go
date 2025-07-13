package errors

// HTTP status error codes as constants
const (
    ErrBadRequest          = 400
    ErrUnauthorized        = 401
    ErrForbidden           = 403
    ErrNotFound            = 404
    ErrMethodNotAllowed    = 405
    ErrConflict            = 409
    ErrUnprocessableEntity = 422
    ErrInternalServerError = 500
    ErrNotImplemented      = 501
    ErrServiceUnavailable  = 503
)

// Standard error response structure
type ErrorResponse struct {
    StatusCode int    `json:"status_code"`
    Error      string `json:"error"`
    Message    string `json:"message"`
}

// Helper to create a standard error response
func NewErrorResponse(code int, err, msg string) ErrorResponse {
    return ErrorResponse{
        StatusCode: code,
        Error:      err,
        Message:    msg,
    }
}
