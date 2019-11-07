package dissball

import (
	"fmt"
	. "fooldealer/src/common"
)

const GAME_ID = Dissball

func init() {
	fmt.Println("disball init route....")
	prefix := "/" + GameMap[GAME_ID] + "/"
	RouteMap[prefix+"ROOM_INFO_REQUEST"] = Demo
}
