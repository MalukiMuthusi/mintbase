package models

type UserIDParameter struct {
	User string `uri:"user" binding:"required"`
}

type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CustomErrorCode int32

const (
	InvalidUserIdParam CustomErrorCode = iota
)

func (e CustomErrorCode) String() string {
	switch e {
	case InvalidUserIdParam:
		return "INVALID_USER_ID_PARAMETER"
	default:
		return "FAILED_TO_PROCESS_REQUEST"
	}
}
