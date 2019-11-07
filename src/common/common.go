package common

//路由map
var RouteMap = make(map[string]func(request BaseMsg) (BaseMsg, error))

//游戏map
const Jungle = 1
const Gobang = 2
const Minesweeping = 3
const Othello = 4
const Dissball = 6

/*gobang(2,"五子棋"),
jungle(1,""),
minesweeping(3,"扫雷大乱斗"),
othello(4,""),
dissball(6,"多人桌球");*/
var GameMap = map[int]string{
	Jungle:       "jungle",
	Gobang:       "gobang",
	Minesweeping: "minesweeping",
	Othello:      "othello",
	Dissball:     "dissball"}

type BaseMsg struct {
	Type   string `json:"type"`
	Extend Extend `json:"extend"`
	Data   string `json:"data"`
}

type Extend struct {
	GId       int    `json:"gId"`
	RoomId    int    `json:"roomId"`
	JcnUserID string `json:"jcnuserid"`
}
