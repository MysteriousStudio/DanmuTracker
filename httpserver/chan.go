package httpserver

import (
	"fmt"
	"net/http"
	"time"
)

type Chans struct {
	err  chan error
	info chan string
}

func (c *Chans) InfoChan(r *http.Request, addtionInfo string) {
	c.info <- (fmt.Sprintf("%s, \t%s, \t%s, \tAddtion info: %s", r.RequestURI, r.RemoteAddr, time.Now().String(), addtionInfo))
}
func (c *Chans) ErrChan(r *http.Request, addtionInfo interface{}) {
	c.err <- fmt.Errorf("%s, \t%s, \t%s, \tAddtion info: %s", r.RequestURI, r.RemoteAddr, time.Now().String(), addtionInfo)
}

func NewChans(info chan string, err chan error) *Chans {
	return &Chans{
		err:  err,
		info: info,
	}
}
