package socket

import (
	"fmt"
	"net"
	"strconv"
)

func ServerStart(ip string, port int) {
	strPort := strconv.Itoa(port)
	ln, err := net.Listen("tcp", ip+":"+strPort)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Printf("%s:%s", "error:", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	return
}
