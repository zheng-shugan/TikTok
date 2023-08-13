package result

type Response struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

const (
	SuccessCode   int32 = 0
	ParamErrCode  int32 = 100
	ServerErrCode int32 = 200
	AuthErrCode   int32 = 300
)

const (
	ParamErrMsg string = "wrong input parameter"
	SuccessMsg  string = "success"
)
