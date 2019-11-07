package socket

import (
	"bytes"
	"fmt"
	"fooldealer/src/dispatcher"
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
	ch := make(chan string)
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
					//待优化
					go sendMessage(ch, message.body)
					go handleMessage(ch, conn)
				}
				lastBytes = preBytes
			}
		}
	}
}

func sendMessage(ch chan string, message string) {
	ch <- message
}

func handleMessage(ch chan string, conn net.Conn) {
	for {
		message := <-ch
		//业务逻辑匹配处理
		res, err := dispatcher.Dispatch(message)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			conn.Write(ConvertToBytes(res))
		}
	}
}
