package Console

import "log"

func Err(v ...interface{}) {
	outPut("err:", v)
}

func outPut(t string, v ...interface{}) {
	log.Println(t, v)
}
