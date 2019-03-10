package main

import (
  "io"
  "net/http"
  "github.com/xiaocuixt/xcvideo/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
  w.WriteHeader(errResp.HttpSC)
  resStr, _ := json.Marshal(&errResp.Error)
  io.WriteString(w, string(resStr))
}

// sc为status code
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
  w.WriteHeader(sc)
  io.WriteString(w, resp)
}















