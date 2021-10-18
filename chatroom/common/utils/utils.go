package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/my/repo/chatroom/common/message"
)

// //write
func WritePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write", err)
		return
	}
	return
}

//封裝一個函數讀取client發來的消息
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	// fmt.Println("等待訊息...")
	_, err = conn.Read(buf[:4])
	if err == io.EOF {
		fmt.Println("斷開魂結")
		return
	}

	if err != nil {
		fmt.Println("conn.Read(buf[:4])", err)
		return
	}
	//根據buf[:4]轉成一個uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read", err)
		return
	}
	//要得內容等於[:pkgLen]為止
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("Unmarshal", err)
		return
	}
	return
}
