package dissball

import . "fool_dealer/src/common"

/**
业务代码
 */
func Demo(request BaseMsg) (BaseMsg, error) {
	request.Data = ":uuuuuuuu"
	request.Type = "TYPE_00234"
	
	return request, nil
}
