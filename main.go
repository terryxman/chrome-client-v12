package main

import (
	"./common"
	"./console"
)

var err error

func main() {
	err = Common.InitConfig()

	if nil != err {
		Console.Err(err.Error())
		return
	}

	Console.Run()
}
