package response

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func Success(c *gin.Context, result Result) {
	Failed(c, result)
}

func Failed(c *gin.Context, result Result) {
	JSON(c, result.ToHTTPCode(), result)
}

var (
	OK                  = newResult(0, "ok")
	InvalidParams       = newResult(10001, "Invalid Parameter")
	Unauthorized        = newResult(10002, "Unauthorized")
	InternalServerError = newResult(10003, "Internal Server Error")
	NotFound            = newResult(10004, "Not Found")
	AlreadyExists       = newResult(10005, "Conflict")
	Timeout             = newResult(10006, "Request Timeout")
	TooManyRequests     = newResult(10007, "Too Many Requests")
	Forbidden           = newResult(10008, "Forbidden")
	LimitExceed         = newResult(10009, "Limit Exceed")
	DeadlineExceeded    = newResult(10010, "Deadline Exceeded")
	AccessDenied        = newResult(10011, "Access Denied")
	MethodNotAllowed    = newResult(10012, "Method Not Allowed")
	ServiceUnavailable  = newResult(10013, "Service Unavailable")
	Canceled            = newResult(10014, "Canceled")
	Unknown             = newResult(10015, "Unknown")
	PermissionDenied    = newResult(10016, "Permission Denied")
	ResourceExhausted   = newResult(10017, "Resource Exhausted")
	FailedPrecondition  = newResult(10018, "Failed Precondition")
	Aborted             = newResult(10019, "Aborted")
	OutOfRange          = newResult(10020, "Out Of Range")
	Unimplemented       = newResult(10021, "Unimplemented")
	DataLoss            = newResult(10022, "Data Loss")
)
