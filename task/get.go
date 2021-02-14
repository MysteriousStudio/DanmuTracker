package task

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type DanmuTask struct {
	PyScript string
}

var pyScript = `
from parsel import Selector

import requests


def get(url):
    headers = {
        "Accept": "*/*",
        "Accept-Language": "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3",
        "User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:55.0) Gecko/20100101 Firefox/55.0"
    }
    body = requests.get(url, headers=headers).content
    xbody = Selector(text=str(body, encoding='utf-8'))
    lists = xbody.xpath("//d")
    count = xbody.xpath("//maxlimit/text()").extract_first()
    for li in lists:
        content = li.xpath("./text()").extract_first()
        par = li.xpath("./@p").extract_first()
        print(content, "&&&", par)


if __name__ == '__main__':
    url = "https://comment.bilibili.com/$CID$.xml"
    get(url)
`

func NewDanmuTask() *DanmuTask {
	return &DanmuTask{
		PyScript: pyScript,
	}
}

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

func (d *DanmuTask) GetDanmu(cid string) (danmu []DanmuContent, err error) {
	var path string

	if runtime.GOOS == "windows" {
		path, err = exec.LookPath("python3.exe")
	} else {
		path, err = exec.LookPath("python3")
	}
	if err != nil {
		return
	}
	cwd, _ := os.Getwd()
	rad := cwd + PathDiv() + strconv.FormatInt(rand.Int63n(1000000), 10) + ".py"
	// fmt.Println("Exec file: " + rad)
	tmpPythonFileContent := strings.ReplaceAll(d.PyScript, "$CID$", cid)
	fp, err := os.Create(rad)
	fp.WriteString(tmpPythonFileContent)
	cmd := exec.Command(path, rad)
	b, err := RunCommand(cmd)
	danmu, err = SplitDanmuContentString(string(b))
	if err != nil {
		return
	}

	defer func() {
		fp.Close()
		os.Remove(rad)
	}()
	return
}

func RunCommand(exec *exec.Cmd) (b []byte, err error) {
	return exec.CombinedOutput()
}

func SplitDanmuContentString(s string) (danmu []DanmuContent, err error) {
	tmpSlice := strings.Split(s, "\n")
	danmu = make([]DanmuContent, 0)

	for _, v := range tmpSlice {
		tmpSlice1 := strings.Split(v, " &&& ")
		if len(tmpSlice1) == 1 {
			// fmt.Println(v)
			continue
		}
		tmpContentSlice := make([]string, 0)
		tmpContentSlice = append(tmpContentSlice, tmpSlice1[0])
		tmpContentSlice = append(tmpContentSlice, strings.Split(tmpSlice1[1], ",")...)

		tmpDanmu := DanmuContent{
			Content:       tmpContentSlice[0],
			AppearTime:    tmpContentSlice[1],
			Mode:          tmpContentSlice[2],
			FontSize:      tmpContentSlice[3],
			Color:         tmpContentSlice[4],
			SendTimeStamp: tmpContentSlice[5],
			DanmuPool:     tmpContentSlice[6],
			DataBaseID:    tmpContentSlice[7],
		}
		danmu = append(danmu, tmpDanmu)
	}
	return
}

func PathDiv() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}
