package Console

import (
	"../common"
	"log"
)

func (a *Action) GetLine() {
	SDataAll, _ = FastGetDataJson(FileList.SSJson)
	Common.ConsoleConfig.Get("get")

}
