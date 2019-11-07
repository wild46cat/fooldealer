package common

//自定义error

const (
	///////////////////////////////////////////////////////////
	ERR_OK               = 0
	ERR_NOGAMEID         = 1
	ERR_NOHANDLEFUNC     = 2
	ERR_BIZHANDLEERROR   = 3
	ERR_JSONCONVERTERROR = 4
	/////////////////////////////////////////////////////////
	unknownMsg = "未知错误"
)

var errorMap = map[int]string{
	ERR_OK:               "ok",
	ERR_NOGAMEID:         "gameId不存在",
	ERR_NOHANDLEFUNC:     "没有对应的处理函数",
	ERR_BIZHANDLEERROR:   "业务处理失败",
	ERR_JSONCONVERTERROR: "json转换失败",
}

func NewError(code int) error {
	msg, ok := errorMap[code]
	if !ok {
		return &baseError{
			ErrorCode:    code,
			ErrorMessage: unknownMsg,
		}
	} else {
		return &baseError{
			ErrorCode:    code,
			ErrorMessage: msg,
		}
	}
}

type baseError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e *baseError) Error() string {
	return e.ErrorMessage
}
