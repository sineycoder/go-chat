package main

import (
	"coding"
	"encoding/json"
	"fmt"
	"message"
	"net"
)

// 写一个函数，完成登录
func login(userId int, userPwd string) (conn net.Conn, err error) {
	// 开始定协议
	conn, err = net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net Dial error")
		return
	}
	// 准备conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	// 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	b, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal error")
		return
	}
	mes.Data = string(b)
	data, err := coding.Encode(mes)
	if err != nil {
		fmt.Println("data Marshal error")
		return
	}
	_, _ = conn.Write(data)
	return
}
