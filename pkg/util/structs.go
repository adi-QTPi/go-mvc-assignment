package util

type StandardResponseJson struct {
	Msg            string `json:"msg"`
	Err            error  `json:"err"`
	ErrDescription string `json:"err_description"`
}
