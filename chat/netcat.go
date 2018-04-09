// 一个连接TCP服务器的简单读写客户端

package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 忽略错误
		log.Println("done")
		done <- struct{}{} // 指示主goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台goroutine完成
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
