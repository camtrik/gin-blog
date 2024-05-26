package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Internal Server Error")
	InvalidParams             = NewError(10000001, "Invalid Parameters")
	NotFound                  = NewError(10000002, "Not Found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Authentication Failed: AppKey and AppSecret Not Found")
	UnauthorizedTokenError    = NewError(10000004, "Authentication Failed: Invalid Token")
	UnauthorizedTokenTimeout  = NewError(10000005, "Authentication Failed: Token Timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Authentication Failed: Token Generation Failed")
	NeedToken                 = NewError(10000007, "Authentication Failed: Token Required. Please generate a token first.")
	TooManyRequests           = NewError(10000008, "Too Many Requests")
)
