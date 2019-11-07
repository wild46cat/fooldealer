package socket

import "testing"

func TestConvertToBytes(t *testing.T) {
	str := "abcedfg89"
	res := ConvertToBytes(str)
	t.Log(res)
}
