package response

const (
	ErrCodeSuccess  = 2001  // Success
	ErrCodeParam    = 20003 // Email invalid
	ErrUnauthorized = 401
	ErrPayloadData  = 400
)

// message
var msg = map[int]string{
	ErrCodeSuccess:  "Success",
	ErrCodeParam:    "Email is invalid",
	ErrUnauthorized: "Unauthorized",
	ErrPayloadData:  "Invalid request payload",
}
