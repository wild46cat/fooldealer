package minesweeping

import (
	. "fooldealer/src/common"
	"testing"
)

func TestDemo(t *testing.T) {
	msg := BaseMsg{
		Type: "0",
		Extend: Extend{
			GId:       0,
			RoomId:    0,
			JcnUserID: "ccc",
		},
		Data: "ljkjkljk",
	}
	res, err := Demo(msg)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(res)
}
