package main

import (
	"fmt"

	"github.com/my/repo/chatroom/client/login"
)

var (
	userId  int
	userPwd string
)

func main() {
	var (
		key  int         //接收用戶選擇
		loop bool = true //是否循環回主畫面
	)
	for loop {
		fmt.Println("=====多人聊天系統=====")
		fmt.Println("\t1.登入")
		fmt.Println("\t2.註冊")
		fmt.Println("\t3.退出")
		fmt.Println("\t請選擇(1-3)...")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("===登入===")
			loop = false
		case 2:
			fmt.Println("===註冊===")
			loop = false
		case 3:
			fmt.Println("你已成功退出")
			loop = false
		default:
			fmt.Println("輸入錯誤，重新輸入")
		}
	}

	if key == 1 {
		fmt.Println("輸入ID...")
		fmt.Scanln(&userId)
		fmt.Println("輸入密碼...")
		fmt.Scanln(&userPwd)
		login.Login(userId, userPwd)
	}

}
