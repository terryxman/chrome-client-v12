package Console

import (
	"../common"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"
)

type actionNameMaps map[string]reflect.Value

type Action struct {
}

type ActionConf struct {
}

func Run() {

	//取得运行配置
	getConfig()
	AutoRun(Common.ConsoleConfig.GetString("run"))

}

func AutoRun(funcName string) {
	var actionFuncs Action
	actionMap := make(actionNameMaps, 0)
	af := reflect.ValueOf(&actionFuncs)
	afType := af.Type()
	afNumber := af.NumMethod()
	for i := 0; i < afNumber; i++ {
		aName := afType.Method(i).Name
		if aName == funcName {
			af.Method(i).Call(nil)
			return
		}
		actionMap[aName] = af.Method(i)
	}
	log.Println("func not exists:", funcName)
	log.Println(Common.FastJsonMarshal(actionMap))

}

//修改自 https://studygolang.com/articles/7706
func _AutoRun(args ...interface{}) {
	var actionFuncs Action
	actionMap := make(actionNameMaps, 0)
	af := reflect.ValueOf(&actionFuncs)
	afType := af.Type()

	afNumber := af.NumMethod()

	for i := 0; i < afNumber; i++ {
		aName := afType.Method(i).Name
		actionMap[aName] = af.Method(i)
	}
	str := "test"
	parms := []reflect.Value{reflect.ValueOf(&str)}
	log.Println(actionMap["Status"])
	actionMap["Status"].Call(parms)

}

func (a *Action) TT(...struct{}) {
	log.Println("fjdkslfjdslkfjsldkfjdslfjlsfjdksljkl")
}

func getConfig() (err error) {
	runPath, err := Common.GetCurrentPath()

	Common.ConsoleConfig.SetConfigName("run")
	Common.ConsoleConfig.SetConfigType("json")
	Common.ConsoleConfig.AddConfigPath(runPath)

	err = Common.ConsoleConfig.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Println(runPath)
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	return

}

func (a *Action) Status() error {
	t := time.Now() //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	log.Println("status：", t.Format("2006-01-02 15:04:05"))
	log.Println("status：")
	return nil
}
