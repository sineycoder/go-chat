package main

import (
	"bufio"
	"coding"
	"encoding/json"
	"fmt"
	"message"
	"net"
	"redisCli"
	"strconv"
)

var r *redisCli.RedisCli

var online map[net.Conn]message.LoginMes

func init() {
	r, _ = redisCli.NewRedisCli("tcp", "127.0.0.1:6379")
	online = make(map[net.Conn]message.LoginMes, 10)
}

func handleLogin(b []byte) (loginMes message.LoginMes) {
	_ = json.Unmarshal(b, &loginMes)
	//使用redis操作是否存在该账号
	str, err := r.GetString(strconv.Itoa(loginMes.UserId) + ":" + loginMes.UserPwd)
	if err != nil {
		return
	}
	//if str == "" {
	//	fmt.Println("登录失败")
	//	return
	//}
	loginMes.UserName = str
	return
}

func handleRegister(b []byte) {

}

func process(conn net.Conn) {
	// 读取客户端发送的信息
	reader := bufio.NewReader(conn)
	mes := message.Message{}
	b, err := coding.Decode(reader)
	if err != nil {
		fmt.Println("decode失败")
		return
	}
	err = json.Unmarshal(b, &mes)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	dataBytes := []byte(mes.Data)
	switch mes.Type {
	case message.LoginMesType: // 登录消息
		loginMes := handleLogin(dataBytes)
		online[conn] = loginMes
		mes := message.Message{
			Type: message.MesType,
			Data: loginMes.UserName + "加入群聊啦",
		}
		b, err := coding.Encode(mes)
		if err != nil {
			fmt.Println("序列化失败")
			return
		}
		for c, _ := range online {
			_, _ = c.Write(b)
		}
	case message.RegisterMesType: // 注册消息
		handleRegister(dataBytes)
	}
}

func main() {
	fmt.Println("服务器从8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen error")
		return
	}
	//监听成功等待
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net client error")
			return
		}
		//连接成功就保持通讯
		go process(conn)
	}
}
