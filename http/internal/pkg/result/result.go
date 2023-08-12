package result

type Response struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

const (
	SuccessCode   int64 = 0
	ParamErrCode  int64 = 100
	ServerErrCode int64 = 200
	AuthErrCode   int64 = 300
)

const (
	ParamErrMsg string = "wrong input parameter"
)
