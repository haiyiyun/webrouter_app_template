package base

import "net/http"

func (self *Service) SomeFunc(rw http.ResponseWriter) {
	rw.Write([]byte("some_func"))
}
