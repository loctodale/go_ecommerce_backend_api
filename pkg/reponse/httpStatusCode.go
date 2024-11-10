package reponse

const (
	CodeSuccess            = 1  // Success
	ErrCodeParamInvalid    = -4 // Email is invalid
	ErrInvalidToken        = -3 // Token invalid
	ErrorCodeUserHasExited = -2 // User is existed
)

// message
var msg = map[int]string{
	CodeSuccess:            "Success",
	ErrCodeParamInvalid:    "Email is invalid",
	ErrInvalidToken:        "Token is invalid",
	ErrorCodeUserHasExited: "User has existed",
}
