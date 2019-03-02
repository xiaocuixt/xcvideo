package main

import (
  "io"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  io.WriteString(w, "Create User Hander")
}

//WriteString第一个参数表示StringWriter，即将第二个参数的内容写入到w中
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  username := ps.ByName("username")
  io.WriteString(w, username)
}