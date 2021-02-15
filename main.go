package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"strconv"

	"github.com/MysteriousStudio/DanmuTracker/conf"
	"github.com/MysteriousStudio/DanmuTracker/httpserver"
	"github.com/kpango/glg"
)

var config conf.TrackerConf
var userHome, _ = os.UserHomeDir()
var infoChan chan string
var errChan chan error

func init() {
	tmpConf := conf.TrackerConf{
		AdminToken: strconv.FormatInt(rand.Int63(), 10),
		ServerPort: "8000",
	}
	b, _ := json.Marshal(tmpConf)
	fp, err := os.Open(userHome + PathDiv() + ".danmutracker_conf.json")
	defer fp.Close()
	if !os.IsExist(err) {
		config = tmpConf
		fp.Close()
		fp, err = os.Create(userHome + PathDiv() + ".danmutracker_conf.json")
		fp.Write(b)
	} else {
		b, _ = ioutil.ReadAll(fp)
		json.Unmarshal(b, &config)
	}
	if err != nil {
		os.Exit(1)
	}
	infoChan = make(chan string)
	errChan = make(chan error)
}

func main() {
	glg.Info("Succeed to init conf!")
	glg.Infof("The admin token is \"%s\", server port is \"%s\"", config.AdminToken, config.ServerPort)
	glg.Info("Starting HTTP server...")
	go httpserver.StartHTTPServer(config.ServerPort, infoChan, errChan)
	glg.Info("Succeed to start DanmuTracker!")

	for {
		select {
		case info := <-infoChan:
			glg.Log(info)
		case err := <-errChan:
			glg.Error(err)
		}
	}
}

// PathDiv ...
func PathDiv() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}
