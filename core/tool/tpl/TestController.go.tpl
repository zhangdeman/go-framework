package controller

func TestMethod() map[string]interface{}  {
	data := make(map[string]string)
	data["name"] = "zhangdeman"
	data["age"] = "22"
	data["high"] ="180"
	returnData := make(map[string]interface{})
	returnData["errCode"] = 200
	returnData["errMsg"] = "success"
	returnData["data"] = data
	return returnData
}