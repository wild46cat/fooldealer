package fooldealer

import "fooldealer/socket"

func main() {
	//socket.ServerStart("localhost", 65500)
	socket.ClientStart("localhost", 65500)
}
