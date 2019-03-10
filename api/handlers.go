package main

import (
  "io"
  "io/ioutil"
  "encoding/json"
  "net/http"
  "log"
  "github.com/xiaocuixt/xcvideo/api/defs"
  "github.com/xiaocuixt/xcvideo/api/dbops"
  "github.com/xiaocuixt/xcvideo/api/session"
  "github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  res, _ := ioutil.ReadAll(r.Body)
  ubody := &defs.UserCredential{}
  log.Printf("%s", r.Body)
  log.Printf("%s", res)
  log.Printf("%s", ubody)
  log.Printf("%s", ps)
  if err := json.Unmarshal(res, ubody); err != nil {
    log.Printf("%s", ubody)
    sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
    return
  }
  if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
    sendErrorResponse(w, defs.ErrorDBError)
    return
  }
  id := session.GenerateNewSessionId(ubody.Username)
  su := &defs.SignedUp{Success: true, SessionId: id}

  if resp, err := json.Marshal(su); err != nil {
    sendErrorResponse(w, defs.ErrorInternalFaults)
    return
  } else {
    sendNormalResponse(w, string(resp), 201)
  }
}

//WriteString第一个参数表示StringWriter，即将第二个参数的内容写入到w中
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  username := ps.ByName("username")
  io.WriteString(w, username)
}