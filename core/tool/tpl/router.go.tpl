package router

import (
	"controller"
)

var RouterMap map[string] func()map[string]interface{}

func init()  {
	RouterMap = make(map[string] func()map[string]interface{})
	RouterMap["/"] = controller.TestRootMethod
	RouterMap["/test"] = controller.TestMethod
}
