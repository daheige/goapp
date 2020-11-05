package ckeys

var (
	// AppName app name
	AppName = CtxKey{"app_name"}

	// XRequestID request_id
	XRequestID = CtxKey{"x-request-id"}

	// ClientIp
	ClientIP = CtxKey{"client_ip"}

	// RequestMethod request method
	RequestMethod = CtxKey{"request_method"}

	// RequestURI request uri
	RequestURI = CtxKey{"request_uri"}

	// UserAgent ua
	UserAgent = CtxKey{"user_agent"}

	// Plat plat
	Plat = CtxKey{"plat"}

	// Detail logger detail key
	Detail = CtxKey{"detail"}
)
