package httpserver

import (
	"net/http"
)

type HandleIndex struct {
	*Handle
}

func (h *HandleIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
