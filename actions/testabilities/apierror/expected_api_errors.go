package apierror

import "fmt"

const (
	MissingAuthHeaderJSON  = `{ "code":"error-unauthorized-auth-header-missing", "message":"missing auth header" }`
	AdminNotAuthorizedJSON = `{ "code":"error-admin-auth-on-user-endpoint", "message":"cannot call user's endpoints with admin authorization" }`
	UserNotAuthorizedJSON  = `{ "code": "error-unauthorized-xpub-not-an-admin-key", "message": "xpub provided is not an admin key" }`
	CannotBindBodyJSON     = `{"code":"error-bind-body-invalid", "message":"cannot bind request body"}`
)

func ExpectedJSON(code, message string) string {
	return fmt.Sprintf(`{"code":"%s","message":"%s"}`, code, message)
}
