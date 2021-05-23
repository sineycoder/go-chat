package main

import (
	"bufio"
	"coding"
	"context"
	"fmt"
	"message"
	"sync"
)

var (
	userId  int
	userPwd string
)

var wg sync.WaitGroup

func main() {
	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("\t\t\t 1 登录聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天系统")
			loop = false
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误 请重新输入")
		}
	}

	if key == 1 {
		// 说明用户要登录
		fmt.Println("请输入用户id")
		_, _ = fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		_, _ = fmt.Scanf("%s\n", &userPwd)
		// 先把登录的函数写到另一个文件中
		conn, err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败了")
		} else {
			fmt.Println("登录成功了")
			// 这里需要监听服务器发来的信息
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			wg.Add(2)
			// 开启接收端进程
			go func(ctx context.Context, cancel context.CancelFunc) {
				fmt.Println("开启接收进程")
				for {
					select {
					case <-ctx.Done():
						fmt.Println("退出接收进程")
						wg.Done()
						return
					default:
						b, err := coding.Decode(bufio.NewReader(conn))
						if err != nil {
							fmt.Println("退出接收进程, err:", err)
							wg.Done()
							return
						}
						str := string(b)
						fmt.Println("收到消息：", str)
						if "exit" == str {
							cancel()
						}
					}
				}
			}(ctx, cancel)
			// 开启发送端进程
			go func(ctx context.Context) {
				fmt.Println("开启写入进程")
				for {
					select {
					case <-ctx.Done():
						fmt.Println("退出写入进程")
						wg.Done()
						return
					default:
						var str string
						_, _ = fmt.Scanf("%s\n", &str)
						mes := message.Message{
							Type: message.MesType,
							Data: str,
						}
						b, err := coding.Encode(mes)
						if err != nil {
							fmt.Println("序列化失败")
							continue
						}
						_, _ = conn.Write(b)
					}
				}
			}(ctx)
			wg.Wait()
		}
	} else if key == 2 {
		fmt.Println("进行用户注册")
	}

}
