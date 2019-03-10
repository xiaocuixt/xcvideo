package main

import (
  "net/http"
  "github.com/xiaocuixt/xcvideo/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_NAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
  sid := r.Header.Get(HEADER_FIELD_SESSION)
  if len(sid) == 0 {
    return false
  }
  uname, ok := session.isSessionExpired(sid)
  if ok {
    return false
  }
  r.Header.Add(HEADER_FIELD_NAME, uname)
  return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
  uname := r.Header.Get(HEADER_FIELD_NAME)
  if len(uname) == 0 {
    sendErrorResponse()
    return false
  }

}