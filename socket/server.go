package socket

import (
	"bytes"
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
	//conn.Write([]byte("hello,client \r\n"))
	defer conn.Close()
	var lastBytes []byte
	for {
		//处理粘包半包问题
		data := bytes.NewBuffer(make([]byte, BUFFER_SIZE))
		n, err := conn.Read(data.Bytes())
		if err != nil {
			break
		} else {
			//模拟延迟情况
			//time.Sleep(time.Duration(3) * time.Second)
			realData := append(lastBytes, data.Bytes()[0:n]...)
			//传递的参数与为有效bytes数据
			messageInfos, preBytes, e := ConvertFromBytes(realData)
			if e != nil {
				fmt.Println(e)
				continue
			} else {
				for _, message := range messageInfos {
					fmt.Println(message.body)
				}
				lastBytes = preBytes
			}
		}
	}
}
