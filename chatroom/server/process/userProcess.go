package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/my/repo/chatroom/common/message"
	"github.com/my/repo/chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

//處理登陸請求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	fmt.Println("處理用戶Login中")
	var loginMes message.LoginMes
	_ = json.Unmarshal([]byte(mes.Data), &loginMes)

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500
	}
	data, _ := json.Marshal(loginResMes)
	resMes.Data = string(data)
	data, _ = json.Marshal(resMes)
	//因為使用了分層模式，先創一個transfer實例然後讀取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	return
}
