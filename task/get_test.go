package task

import (
	"testing"
)

var dt *DanmuTask = NewDanmuTask()

var CID string = "268620037"

func TestGetCID(t *testing.T) {
	checkFlag := false
	targetCID := make([]string, 0)
	targetCID = append(targetCID, "268620037")
	cid, err := dt.GetCID("BV1Va411c7")
	for k, v := range cid {
		t.Logf("[%d]%s", k, v)
	}
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	}
	for k, v := range cid {
		if v == targetCID[k] {
			checkFlag = true
			return
		}
	}
	if checkFlag != true {
		t.Fail()
		t.Log("No CID matched!")
		t.FailNow()
	}
}

func TestGetDanmu(t *testing.T) {
	d, err := dt.GetDanmu(CID)
	if err != nil {
		t.Fail()
		t.Log(err.Error())
		t.Log(d)
	}
}
