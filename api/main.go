package main

import (
   "net/http"
   "github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
  r *httprouter.Router
}

func newMiddleWareHandler(r *httprouter) http.Handler {
  m := middleWareHandler{}
  m.r = r
  return m
}

func (m middleWareHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
  // check session
  validateUserSession(r)
  m.r.serveHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
  router := httprouter.New()
  router.POST("/user", CreateUser)
  router.POST("/user/:username", Login)
  return router
}

func main(){
  r := RegisterHandlers()
  mh := newMiddleWareHandler(r)
  http.ListenAndServe(":8000", mh)
}