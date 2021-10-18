package processor

import (
	"fmt"
	"io"
	"net"

	"github.com/my/repo/chatroom/common/message"
	"github.com/my/repo/chatroom/server/process"
	"github.com/my/repo/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//根據client發送的消息決定調用哪種函數
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//處理登陸
		//創建一個userProcess實例
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息類型不存在")
	}
	return
}

func (this *Processor) Process2() (err error) {
	for {
		//創一個transfer
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err == io.EOF {
			fmt.Println("斷開魂結")
			return err
		}
		if err != nil {
			fmt.Println("readPkg err", err)
			return err
		}
		fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
