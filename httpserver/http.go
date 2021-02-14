package httpserver

import "net/http"

type Handle struct {
	InfoChan chan string
	ErrChan  chan error
}

func StartHTTPServer(port string, infoChan chan string, errChan chan error) (err error) {
	handle := &Handle{
		InfoChan: infoChan,
		ErrChan:  errChan,
	}
	http.Handle("/", &HandleIndex{Handle: handle})
	http.Handle("/index", &HandleIndex{Handle: handle})
	http.Handle("/get", &HandleGet{Handle: handle})
	err = http.ListenAndServe("127.0.0.1:"+port, nil)
	errChan <- err
	return
}
