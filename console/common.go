package Console

import (
	"../common"
	"io/ioutil"
)

func FastPutData(fileName string, data []byte) error {
	var fullName = "data/" + fileName
	return ioutil.WriteFile(fullName, data, 07777)
}

func FastGetData(fileName string) (data []byte, err error) {
	var fullName = "data/" + fileName
	data, err = ioutil.ReadFile(fullName)
	return
}

func FastGetDataJson(fileName string) (data []SData, err error) {
	var _d []byte
	_d, err = FastGetData(fileName)
	if nil != err {
		return nil, err
	}
	err = Common.Json.Unmarshal(_d, &data)

	return
}
