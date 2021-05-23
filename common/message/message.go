package message

const (
	LoginMesType    = "LoginMes"
	RegisterMesType = "RegisterMesType"
	MesType         = "MesType"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type loginRes struct {
	Code  int    `json:"code"`  // 状态码
	Error string `json:"error"` // 错误信息
}
