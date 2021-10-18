package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/my/repo/chatroom/common/message"
	"github.com/my/repo/chatroom/common/utils"
)

func Login(userId int, userPwd string) (err error) {
	//定協議
	fmt.Printf("輸入的ID=%d 密碼=%v", userId, userPwd)

	//連接到服務器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial", err)
	}
	defer conn.Close()

	//準備發消息給服務器
	var mes message.Message
	mes.Type = message.LoginMesType
	//創建一個loginMes結構體
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//將loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("Marshal(loginMes)", err)
		return
	}
	//把data賦給mes.Data字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes)", err)
		return
	}
	//到現在data就是發出去的包
	//先獲取data的長度 轉成一個切片
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	fmt.Println("客戶端發送長度=", len(data))
	//發送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	mes, err = utils.ReadPkg(conn) //mes
	if err != nil {
		fmt.Println("ReadPkg(conn)", err)
		return
	}

	//接收返回的登入驗證訊息
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登錄成功")
	} else if loginResMes.Code == 500 {
		fmt.Println("登錄失敗", loginResMes.Error)
	}
	return
}
