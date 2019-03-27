package defs


type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorReseponse struct {
	HttpSC int
	Error Err
}

var(
	ErrorRequestBodyParseFaild = ErrorReseponse{HttpSC:400,Error:Err{Error:"Request body is not correct",ErrorCode:"001"}}
	ErrorNotAuthUser = ErrorReseponse{HttpSC:401,Error:Err{Error:"User authentication failed",ErrorCode:"002"}}
	ErrorDBError = ErrorReseponse{HttpSC:500,Error:Err{Error:"DB ops faild",ErrorCode:"003"}}
	ErrorInterfnalFault = ErrorReseponse{HttpSC:500,Error:Err{Error:"Internal service error",ErrorCode:"004"}}

	)