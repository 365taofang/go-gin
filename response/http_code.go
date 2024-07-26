package response

type CodeS struct {
	Success             *ErrorCode
	InvalidParams       *ErrorCode
	Unauthorized        *ErrorCode
	InternalServerError *ErrorCode
	NotFound            *ErrorCode
	AlreadyExists       *ErrorCode
	Timeout             *ErrorCode
	TooManyRequests     *ErrorCode
	Forbidden           *ErrorCode
	LimitExceed         *ErrorCode
	DeadlineExceeded    *ErrorCode
	AccessDenied        *ErrorCode
	MethodNotAllowed    *ErrorCode
	ServiceUnavailable  *ErrorCode
	Canceled            *ErrorCode
	Unknown             *ErrorCode
	PermissionDenied    *ErrorCode
	ResourceExhausted   *ErrorCode
	FailedPrecondition  *ErrorCode
	Aborted             *ErrorCode
	OutOfRange          *ErrorCode
	Unimplemented       *ErrorCode
	DataLoss            *ErrorCode
}

var Code = &CodeS{
	Success:             NewError(0, "ok"),
	InvalidParams:       NewError(10001, "Invalid Parameter"),
	Unauthorized:        NewError(10002, "Unauthorized"),
	InternalServerError: NewError(10003, "Internal Server Error"),
	NotFound:            NewError(10004, "Not Found"),
	AlreadyExists:       NewError(10005, "Conflict"),
	Timeout:             NewError(10006, "Request Timeout"),
	TooManyRequests:     NewError(10007, "Too Many Requests"),
	Forbidden:           NewError(10008, "Forbidden"),
	LimitExceed:         NewError(10009, "Limit Exceed"),
	DeadlineExceeded:    NewError(10010, "Deadline Exceeded"),
	AccessDenied:        NewError(10011, "Access Denied"),
	MethodNotAllowed:    NewError(10012, "Method Not Allowed"),
	ServiceUnavailable:  NewError(10013, "Service Unavailable"),
	Canceled:            NewError(10014, "Canceled"),
	Unknown:             NewError(10015, "Unknown"),
	PermissionDenied:    NewError(10016, "Permission Denied"),
	ResourceExhausted:   NewError(10017, "Resource Exhausted"),
	FailedPrecondition:  NewError(10018, "Failed Precondition"),
	Aborted:             NewError(10019, "Aborted"),
	OutOfRange:          NewError(10020, "Out Of Range"),
	Unimplemented:       NewError(10021, "Unimplemented"),
	DataLoss:            NewError(10022, "Data Loss"),
}
