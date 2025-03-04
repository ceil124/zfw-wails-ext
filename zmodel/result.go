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

func Success(msg string) ResultVO {
	return ResultVO{
		Code: 200,
		Msg:  msg,
	}
}

func SuccessWithData(msg string, data string) ResultVO {
	return ResultVO{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

func SuccessOk() ResultVO {
	return Success(OK)
}

func SuccessOkWithData(data string) ResultVO {
	return SuccessWithData(OK, data)
}

func Failure(msg string) ResultVO {
	return ResultVO{
		Code: 500,
		Msg:  msg,
	}
}

func FailureWithData(msg string, data string) ResultVO {
	return ResultVO{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}
