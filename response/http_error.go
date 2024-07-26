package response

import (
	"fmt"
	"net/http"
)

type ErrorCode struct {
	code int
	msg  string
}

func NewError(code int, msg string) *ErrorCode {
	return &ErrorCode{code: code, msg: msg}
}

func (e *ErrorCode) Code() int {
	return e.code
}

func (e *ErrorCode) Msg() string {
	return e.msg
}

func (e *ErrorCode) Err() error {
	return fmt.Errorf("code = %d, msg = %s", e.code, e.msg)
}

func (e *ErrorCode) ToHTTPCode() int {
	switch e.Code() {
	case Code.Success.Code():
		return http.StatusOK
	case Code.InternalServerError.Code():
		return http.StatusInternalServerError
	case Code.InvalidParams.Code():
		return http.StatusBadRequest
	case Code.Unauthorized.Code(), Code.PermissionDenied.Code():
		return http.StatusUnauthorized
	case Code.TooManyRequests.Code(), Code.LimitExceed.Code():
		return http.StatusTooManyRequests
	case Code.Forbidden.Code(), Code.AccessDenied.Code():
		return http.StatusForbidden
	case Code.NotFound.Code():
		return http.StatusNotFound
	case Code.AlreadyExists.Code():
		return http.StatusConflict
	case Code.Timeout.Code(), Code.DeadlineExceeded.Code():
		return http.StatusRequestTimeout
	case Code.MethodNotAllowed.Code():
		return http.StatusMethodNotAllowed
	case Code.ServiceUnavailable.Code():
		return http.StatusServiceUnavailable
	case Code.Unimplemented.Code():
		return http.StatusNotImplemented
	}

	return http.StatusInternalServerError
}
