package coding

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func Encode(i interface{}) (b []byte, err error) {
	//将message进行序列化
	data, err := json.Marshal(i)
	if err != nil {
		fmt.Println("json Marshal error")
		return
	}
	//这个时候data就是要发送的消息
	var pkgLen uint32
	var pkg = new(bytes.Buffer)
	pkgLen = uint32(len(data))
	// 写入头
	err = binary.Write(pkg, binary.LittleEndian, pkgLen)
	// 写入体
	err = binary.Write(pkg, binary.LittleEndian, data)
	b = pkg.Bytes()
	return
}

func Decode(reader *bufio.Reader) (b []byte, err error) {
	// 读取消息
	preBytes, err := reader.Peek(4)
	var length uint32
	err = binary.Read(bytes.NewBuffer(preBytes), binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("读取头4字节失败", err)
		return
	}
	if uint32(reader.Buffered()) < length+4 {
		return []byte{}, fmt.Errorf("读取到%d，应为%d", reader.Buffered(), length+4)
	}
	fmt.Println("body长度为", length)
	p := make([]byte, length+4)
	_, err = reader.Read(p)
	if err != nil {
		fmt.Println("读取头数据体失败")
		return
	}
	b = p[4:]
	return
}
