package main

import (
	"fmt"
	"net"

	"github.com/my/repo/chatroom/server/main/processor"
)

func processM(conn net.Conn) {
	//調用總控
	processor := &processor.Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("協程錯誤", err)
		return
	}

}
func main() {
	//提示
	fmt.Println("伺服在8889端口監聽")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		fmt.Println("等待客戶端連線中...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("", err)
			return
		}
		go processM(conn)
	}
}
