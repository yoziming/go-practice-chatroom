package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/my/repo/chatroom/common/message"
)

//將方法關連到結構體中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //傳輸時使用的緩存
}

//封裝一個函數讀取client發來的消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	// fmt.Println("等待訊息...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err == io.EOF {
		fmt.Println("斷開魂結")
		return
	}

	if err != nil {
		fmt.Println("conn.Read(buf[:4])", err)
		return
	}
	//根據buf[:4]轉成一個uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read", err)
		return
	}
	//要得內容等於[:pkgLen]為止
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("Unmarshal", err)
		return
	}
	return
}

//write
func (t *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32 = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	_, err = t.Conn.Write(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	_, err = t.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	return
}
