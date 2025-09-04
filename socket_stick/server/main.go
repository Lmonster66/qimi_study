package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"proto"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() // 建立连接
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("Accept connection from:", conn.RemoteAddr())
		go process(conn) // 启动一个goroutine处理连接
	}
}
