package main

import (
	"fooldealer/src/dispatcher"
	_ "fooldealer/src/service/dissball"
	_ "fooldealer/src/service/minesweeping"
	"fooldealer/src/socket"
	"sync"
)

func main() {
	dispatcher.ShowRoute()
	go socket.ServerStart("localhost", 65500)
	go socket.ClientStart("localhost", 65500)
	holdTheWorld()
}

func holdTheWorld() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
