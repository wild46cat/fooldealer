package minesweeping

import . "fooldealer/src/common"

/**
业务代码
 */
func Demo(request BaseMsg) (BaseMsg, error) {
	request.Data = ":.........."
	request.Type = "TYPE_2343"
	request.Extend.JcnUserID = "aabccdd"
	return request, nil
}
