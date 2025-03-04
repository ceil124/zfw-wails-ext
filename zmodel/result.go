package zmodel

const (
	OK      = "操作成功"
	FAILURE = "操作失败"
)

type ResultVO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SuccessWithData(msg string, data any) ResultVO {
	return ResultVO{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

func Success(msg string) ResultVO {
	return SuccessWithData(msg, nil)
}

func FailureWithData(msg string, data any) ResultVO {
	return ResultVO{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}

func Failure(msg string) ResultVO {
	return FailureWithData(msg, nil)
}
