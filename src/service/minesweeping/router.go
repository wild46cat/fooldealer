package minesweeping

import (
	"fmt"
	. "fooldealer/src/common"
)

const GAME_ID = Minesweeping

func init() {
	fmt.Println("minesweeping init route....")
	prefix := "/" + GameMap[GAME_ID] + "/"
	RouteMap[prefix+"ROOM_INFO_REQUEST"] = Demo
}
