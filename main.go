package main

import (
	"fmt"
	"github.com/yaochen2019/gogofly/util"
)

// @title Go-web开发记录
// @version 0.0.1
// @description 学习golang开发记录
func main() {
	//defer cmd.Clean()
	//cmd.Start()
	token, _ := util.GenerateToken(1, "zs")
	fmt.Print(token)
	iJwt, _ := util.ParseToken(token)
	fmt.Print(iJwt)
}
