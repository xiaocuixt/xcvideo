package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

// middle ware的实现方式
type middleWareHandler struct {
  r *httprouter.Router
  l *ConnLimiter
}

func newMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
  m := middleWareHandler{}
  m.r = r
  m.l = NewConnLimiter(cc)
  return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if !m.l.GetConn(){
    sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
    return
  }
  m.r.ServeHTTP(w, r)
  defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
  router := httprouter.New()
  router.GET("/videos/:vid-id", streamHandler)
  router.POST("/upload/:vid-id", uploadHandler)
  router.GET("/testpage", testPageHandler)

  return router
}

func main() {
  r := RegisterHandlers()
  mh := newMiddleWareHandler(r, 2)  //暂时设置流控值为2
  http.ListenAndServe(":9000", mh)
}
