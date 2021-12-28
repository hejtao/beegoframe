package shared

const (
	// http code
	Hc200 = 200
	Hc400 = 400
	Hc401 = 401
	Hc404 = 404
	Hc500 = 500

	// response code
	RcSuccess          = 20000
	RcActionError      = 30000
	RcParameterError   = 30003
	RcWrongPassword    = 30001
	RcResourceNotExist = 30004
)

const (
	// response message
	RmAccountIsDeleted = "account is deleted"

	TimeLayout      = "2006-01-02 15:04:05"
	TimeLayoutZone  = "2006-01-02T15:04:05Z"
	DateLayout      = "2006-01-02"
	TimeLayoutSlash = "2006/1/2 15:04:05"
	DateLayoutSlash = "2006/1/2"
)
