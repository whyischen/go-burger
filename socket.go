package main

import "net"

func listen() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	//结束时关闭连接
	defer conn.Close()
	//读取连接上的数据
	var buf [1024]byte
	_, _ = conn.Read(buf[:])
	//发送数据
	_, _ = conn.Write([]byte("I am server!"))
}
