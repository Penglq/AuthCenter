package conf

type Error interface {
	GetErrData() *Response
	GetSuccData(interface{}) *Response
}

var ResMap = &resCode{
	SYSTERM_ERROR: "系统错误",
	SUCCESS:       "成功",
	SYSTERM_NONE:  "未知错误",
	AUTH_NODE:     "没有操作权限",
}

type Response struct {
	ResCode string      `json:"code"`
	ResDesc string      `json:"resDesc"`
	Data    interface{} `json:"data"`
}

type resCode map[string]string

func (e *resCode) GetErrData(code string) (res *Response) {
	if v, ok := (*e)[code]; ok {
		res.ResCode = code
		res.ResDesc = v
	} else {
		res.ResCode = SYSTERM_NONE
		res.ResDesc = (*e)[SYSTERM_NONE]
	}
	res.Data = struct{}{}
	return
}

func (e *resCode) GetSuccData(data interface{}) (res *Response) {
	res.ResCode = SUCCESS
	res.ResDesc = (*e)[SUCCESS]
	res.Data = data
	return
}
