package utils

import (
	"Go_Learn/CommunicationSystem/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// 读取数据包
func (this *Transfer) ReadPkg() (meg common.Message, err error) {
	conn := this.Conn
	buf := make([]byte, 8096) // 创建一个字节切片用来存储接收到的数据
	fmt.Println("读取服务端 %s 发来的数据\n", conn.RemoteAddr().String())
	_, err = conn.Read(buf[:4])
	if err != nil {
		err = errors.New("数据读取错误！")
		return
	}
	// 将前4字节转为uint类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// 根据长度读取消息内容
	n, err := conn.Read(buf[:pkgLen]) // 从网络读取pkgLen长度的数据，并覆盖前4个位置
	if n != int(pkgLen) || err != nil {
		err = errors.New("数据读取错误！")
		return
	}
	// 将消息反序列化为Message类型
	err = json.Unmarshal(buf[:pkgLen], &meg)
	if err != nil {
		err = errors.New("反序列化错误！")
		return
	}
	return
}

// 发送数据包
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先把data的长度发给服务器
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	fmt.Printf("发送消息长度为%d 内容为：%s", len(data), string(data))

	// 发送消息本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	return
}
