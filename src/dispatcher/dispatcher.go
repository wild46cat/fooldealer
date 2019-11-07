package dispatcher

import (
	"encoding/json"
	"fmt"
	. "fooldealer/src/common"
)

const ERROR_STR = "{\"type\":\"\",\"extend\":{\"gId\":0,\"roomId\":0,\"jcnuserid\":\"\"},\"data\":\"\"}"

func Dispatch(message string) (string, error) {
	fmt.Println(message)
	//处理编解码
	baseMsg, err := DecodeJsonToBaseMsg(message)
	if err != nil {
		fmt.Println(err)
	}
	//判断gameId是否存在
	gameName, exists := GameMap[baseMsg.Extend.GId]
	if !exists {
		fmt.Println("error gameId.")
		return ERROR_STR, NewError(ERR_NOGAMEID)
	}
	routeKey := "/" + gameName + "/" + baseMsg.Type
	f, ok := RouteMap[routeKey]
	if !ok {
		fmt.Println("no method handle" + routeKey)
		return ERROR_STR, NewError(ERR_NOHANDLEFUNC)
	} else {
		//方法调用(待推敲，能否使用interface)
		resBaseMsg, bizErr := funcConvert(f, baseMsg)
		if bizErr != nil {
			fmt.Print("bizErr:")
			fmt.Println(bizErr)
			return ERROR_STR, NewError(ERR_BIZHANDLEERROR)
		}
		resStr, err := EncodeBaseMsgToJson(resBaseMsg)
		if err != nil {
			fmt.Println("error convert to json string")
			return ERROR_STR, NewError(ERR_JSONCONVERTERROR)
		}
		return resStr, nil
	}
}

func funcConvert(f func(msg BaseMsg) (BaseMsg, error), param BaseMsg) (BaseMsg, error) {
	return f(param)
}

func ShowRoute() {
	for k := range RouteMap {
		fmt.Println("...route:--->" + k)
	}
}

func GenerateTestJson(uri string, gid int) string {
	c := Extend{gid, 0, ""}
	msg := BaseMsg{uri, c, ""}
	result, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(result)
	return jsonStr
}

//json ->BaseMsg
func DecodeJsonToBaseMsg(jsonStr string) (BaseMsg, error) {
	var msg BaseMsg
	err := json.Unmarshal([]byte(jsonStr), &msg)
	return msg, err
}

//BaseMsg ->json
func EncodeBaseMsgToJson(baseMsg BaseMsg) (string, error) {
	res, err := json.Marshal(&baseMsg)
	return string(res), err
}
