package Console

import (
	"fmt"
)

func (a *Action) PrintAll() {
	SDataAll, _ = FastGetDataJson(FileList.SSJson)
	for _, sData := range SDataAll {
		var out = sData.DateString + "  "
		for i := 1; i <= 33; i++ {

			var meet = false
			for _, _num := range sData.R {
				if _num == i {
					meet = true
				}
			}
			if true == meet {
				out = out + "X"
			} else {
				out = out + "-"
			}
		}
		out = out + " | "
		for i := 1; i <= 16; i++ {

			var meet = false
			for _, _num := range sData.B {
				if _num == i {
					meet = true
				}
			}
			if true == meet {
				out = out + "X"
			} else {
				out = out + "-"
			}
		}
		fmt.Println(out)
	}
}
