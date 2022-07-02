package controller

// Code msg

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidVideoURL
	CodeInvalidTime
	CodeNeedLogin
	CodeEditError
	CodeServerError
)

// 存储 Code 及状态描述
var CodeMsgMap = map[ResCode]string{
	CodeSuccess:         "剪辑成功",
	CodeInvalidVideoURL: "请输入正确的视频地址",
	CodeInvalidTime:     "请输入正确的时间参数",
	CodeNeedLogin:       "请登录！",
	CodeEditError:       "剪辑失败,请重试",
	CodeServerError:     "系统错误",
}

// Msg     根据 Code 返回对应的状态描述
func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerError]
	}
	return msg
}
