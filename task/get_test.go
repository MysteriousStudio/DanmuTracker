package task

import (
	"testing"
)

var dt *DanmuTask = NewDanmuTask()

var CID string = "268620037"

var testDanmu = `不用了谢谢 &&& 34.53500,5,25,15138834,1608553684,0,cb79f194,42683077467570179
如果我说没有呢(doge) &&& 4.05100,5,25,15138834,1608553624,0,cb79f194,42683046074253317
就是 &&& 14.78200,1,25,16777215,1608553526,0,85a0bfba,42682994833489923
所有图发我一遍，币给你 &&& 152.17700,1,25,16777215,1608552247,0,fc80677b,42682324032684035
烤只傻狍子 &&& 98.21700,1,25,16777215,1608552176,0,fc80677b,42682286928822275`

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

func TestSplitDanmuContentString(t *testing.T) {
	c, _ := SplitDanmuContentString(testDanmu)
	for k, v := range c {
		t.Logf("[%d]%s", k, v.Content)
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
