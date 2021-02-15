package task

import (
	"compress/flate"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

// DanmuTask ...
type DanmuTask struct {
}

// NewDanmuTask ...
func NewDanmuTask() *DanmuTask {
	return &DanmuTask{}
}

// GetCID ...
func (d *DanmuTask) GetCID(bid string) (cid []string, err error) {
	playInfo := PlayInfo{}
	resp, err := http.Get("https://api.bilibili.com/x/player/pagelist?bvid=" + bid + "os&jsonp=jsonp")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &playInfo)
	for _, v := range playInfo.Data {
		cid = append(cid, strconv.FormatInt(v.Cid, 10))
	}
	return
}

// GetDanmu ...
func (d *DanmuTask) GetDanmu(cid string) (danmu []DanmuContent, err error) {

	danmu, err = SplitDanmuContentString(d.DoGet(cid))
	if err != nil {
		return
	}

	return
}

// SplitDanmuContentString ...
func SplitDanmuContentString(i I) (danmu []DanmuContent, err error) {

	danmu = make([]DanmuContent, 0)

	for _, v := range i.D {

		tmpContentSlice := make([]string, 0)
		tmpContentSlice = append(tmpContentSlice, v.Text)
		tmpContentSlice = append(tmpContentSlice, strings.Split(v.P, ",")...)

		tmpDanmu := DanmuContent{
			Content:       tmpContentSlice[0],
			AppearTime:    tmpContentSlice[1],
			Mode:          tmpContentSlice[2],
			FontSize:      tmpContentSlice[3],
			Color:         tmpContentSlice[4],
			SendTimeStamp: tmpContentSlice[5],
			DanmuPool:     tmpContentSlice[6],
			UserHash:      tmpContentSlice[7],
			DataBaseID:    tmpContentSlice[8],
		}
		danmu = append(danmu, tmpDanmu)
	}
	return
}

// PathDiv ...
func PathDiv() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

func (t *DanmuTask) DoGet(cid string) I {
	url := "https://comment.bilibili.com/%CID%.xml"
	url = strings.ReplaceAll(url, "%CID%", cid)
	client := http.Client{}
	httpReq, _ := http.NewRequest("GET", url, nil)

	httpReq.Header.Set("Accept-Encoding", "gzip,deflate,br")
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0")
	httpReq.Header.Set("Accept-Language", "zh-CN,zh;q=0.5")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	resp, err := client.Do(httpReq)
	if err != nil {
		panic(err)
	}

	//Flate data
	i := flate.NewReader(resp.Body)
	b, err := ioutil.ReadAll(i)
	if err != nil {
		panic(err)
	}

	tmp := I{}
	if err = xml.Unmarshal(b, &tmp); err != nil {
		panic(err)
	}

	return tmp
}
