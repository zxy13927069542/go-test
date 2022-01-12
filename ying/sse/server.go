package sse

import (
	"github.com/alexandrevicenzi/go-sse"
	"net/http"
)

func NewServer() {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	http.Handle("/test/", s)
}
