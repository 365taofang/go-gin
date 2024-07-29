package response

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func newResult(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (r *Result) WithMsg(msg string) Result {
	return Result{
		Code: r.Code,
		Msg:  msg,
		Data: r.Data,
	}
}

func (r *Result) Result() Result {
	return Result{
		Code: r.Code,
		Msg:  r.Msg,
		Data: r.Data,
	}
}

func (r *Result) WithData(data any) Result {
	return Result{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}

func (r *Result) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: r.Code,
		Msg:  r.Msg,
		Data: r.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

func (r *Result) StatusCode() int {
	switch r.Code {
	case OK.Code:
		return http.StatusOK
	case InternalServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case Unauthorized.Code, PermissionDenied.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code, LimitExceed.Code:
		return http.StatusTooManyRequests
	case Forbidden.Code, AccessDenied.Code:
		return http.StatusForbidden
	case NotFound.Code:
		return http.StatusNotFound
	case AlreadyExists.Code:
		return http.StatusConflict
	case Timeout.Code, DeadlineExceeded.Code:
		return http.StatusRequestTimeout
	case MethodNotAllowed.Code:
		return http.StatusMethodNotAllowed
	case ServiceUnavailable.Code:
		return http.StatusServiceUnavailable
	case Unimplemented.Code:
		return http.StatusNotImplemented
	}

	return http.StatusInternalServerError
}
