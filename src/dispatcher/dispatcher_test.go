package dispatcher

import (
	_ "fooldealer/src/service/dissball"
	_ "fooldealer/src/service/minesweeping"
	"testing"
)

func TestDispatch(t *testing.T) {
	jsonStr := GenerateTestJson("ROOM_INFO_REQUESTa", 6)
	dispatch, err := Dispatch(jsonStr)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(dispatch)
	}
}
