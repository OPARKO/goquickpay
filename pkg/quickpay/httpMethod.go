package quickpay

type HTTPMethod string

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods
const (
	Get     HTTPMethod = "GET"
	Head    HTTPMethod = "HEAD"
	Post    HTTPMethod = "POST"
	Put     HTTPMethod = "PUT"
	Delete  HTTPMethod = "DELETE"
	Connect HTTPMethod = "CONNECT"
	Option  HTTPMethod = "OPTION"
	Trace   HTTPMethod = "TRACE"
	Patch   HTTPMethod = "PATCH"
)
