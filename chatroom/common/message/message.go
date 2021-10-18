package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes "
)

//在sc之間傳的包
type Message struct {
	Type string `json:"type"` //訊息類型
	Data string `json:"data"`
}

//從client發給server的登入訊息
type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

//從server發給client發給的登入返回訊息
type LoginResMes struct {
	Code  int    `json:"code"` //錯誤代碼
	Error string `json:"error"`
}

type RegisterMes struct{}
