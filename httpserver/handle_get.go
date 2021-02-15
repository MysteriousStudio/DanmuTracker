package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/MysteriousStudio/DanmuTracker/task"
)

// HandleGet ...
type HandleGet struct {
	*Handle
}

// ServeHTTP ...
func (h *HandleGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	chans := NewChans(h.InfoChan, h.ErrChan)
	query := r.URL.Query()
	bid := query.Get("bid")
	if len(query) == 0 || len(bid) == 0 {
		chans.InfoChan(r, "empty request!")
		return
	} else if len(query.Get("panic")) != 0 {
		defer catchPanic(chans, r)
		panic("a test for panic")
	}

	defer catchPanic(chans, r)

	if b, err := h.getDanmu(bid); err != nil {
		chans.ErrChan(r, err)
	} else {
		if _, err := w.Write(b); err != nil {
			panic(err)
		}
		chans.InfoChan(r, "BID is "+bid)
	}

	chans.InfoChan(r, "success")
	defer catchPanic(chans, r)
}

func catchPanic(chans *Chans, r *http.Request) {
	if e := recover(); e != nil {
		chans.ErrChan(r, e)
	}
}

func (h *HandleGet) getDanmu(BID string) (b []byte, err error) {
	tmpSlice := make([]task.DanmuContent, 0)
	t := task.NewDanmuTask()
	cid, err := t.GetCID(BID)
	if err != nil {
		return
	}
	for _, v := range cid {
		if danmu, err := t.GetDanmu(v); err != nil {
			return nil, err
		} else {
			tmpSlice = append(tmpSlice, danmu...)
		}
	}
	b, err = json.Marshal(tmpSlice)
	return
}
