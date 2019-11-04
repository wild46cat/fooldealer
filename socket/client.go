package socket

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ClientStart(ip string, port int) {
	strPort := strconv.Itoa(port)
	conn, err := net.Dial("tcp", ip+":"+strPort)
	if err != nil {
		// handle error
		fmt.Println(err)
	} else {
		clientHandle(conn)
	}
}

func clientHandle(conn net.Conn) {
	str := "aaabbbcccdddeeeff"
	defer conn.Close()
	fmt.Println("client begin handle")
	for {
		fmt.Println("reading....")
		time.Sleep(time.Duration(1) * time.Second)
		conn.Write(ConvertToBytes(str))
		//recv_data := make([]byte, 1024)
		//n, err := bufio.NewReader(conn).Read(recv_data)
		//if n == 0 || err != nil {
		//	fmt.Println(err)
		//	continue
		//} else {
		//	fmt.Println(string(recv_data[0:n]))
		//	conn.Write([]byte("a"))
		//}
	}
}
