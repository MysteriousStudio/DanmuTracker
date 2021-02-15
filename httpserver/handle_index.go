package httpserver

import (
	"net/http"
)

// HandleIndex ...
type HandleIndex struct {
	*Handle
}

// ServeHTTP ...
func (h *HandleIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
