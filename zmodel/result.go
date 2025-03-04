package zmodel

const (
	OK      = "操作成功"
	FAILURE = "操作失败"
)

type ResultVO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func SuccessWithData(msg string, data string) ResultVO {
	return ResultVO{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

func Success(msg string) ResultVO {
	return SuccessWithData(msg, "")
}

func SuccessWithEmptyList(msg string, data string) ResultVO {
	return SuccessWithData(msg, "[]")
}

func SuccessWithEmptyObject(msg string, data string) ResultVO {
	return SuccessWithData(msg, "{}")
}

func FailureWithData(msg string, data string) ResultVO {
	return ResultVO{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}

func Failure(msg string) ResultVO {
	return FailureWithData(msg, "")
}

func FailureWithEmptyList(msg string) ResultVO {
	return FailureWithData(msg, "[]")
}

func FailureWithEmptyObject(msg string) ResultVO {
	return FailureWithData(msg, "{}")
}
