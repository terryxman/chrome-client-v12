package Console

import "time"

type SData struct {
	ID         string
	Date       time.Time
	DateString string
	R          []int
	B          []int
}

var SDataAll []SData

var FileList = struct {
	SSJson string
}{
	SSJson: "ssjson.json",
}
