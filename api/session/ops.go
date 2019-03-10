package session

import (
  "time"
  "sync"
  "github.com/xiaocuixt/xcvideo/api/defs"
  "github.com/xiaocuixt/xcvideo/api/dbops"
  "github.com/xiaocuixt/xcvideo/api/utils"
)

var sessionMap *sync.Map

func init() {
  sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
  return time.Now().UnixNano()/100000 //毫秒
}

func deleteExpiredSeesion(sid string) {
  sessionMap.Delete(sid)
  dbops.DeleteSession(sid)
}

// 内部操作，没有返回值
func LoadSessionsFromDB() {
  r, err := dbops.RetrieveAllSessions()
  if err != nil {
    return
  }
  r.Range(func(k, v interface{}) bool{
    ss := v.(*defs.SimpleSession)
    sessionMap.Store(k, ss)
    return true
  })
}

func GenerateNewSessionId(username string) string {
  id, _ := utils.NewUUID()
  ctime := nowInMilli()
  ttl := ctime + 30 * 60 * 1000 //session在本地的过期时间为30 minutes

  ss := &defs.SimpleSession{Username: username, TTL: ttl}
  sessionMap.Store(id, ss)
  dbops.InsertSession(id, ttl, username)
  return id
}

func IsSessionExpired(sid string) (string, bool) {
  ss, ok := sessionMap.Load(sid)
  if ok {
    ct := nowInMilli()
    if ss.(*defs.SimpleSession).TTL < ct {
      deleteExpiredSeesion(sid)
      return "", true
    }
    return ss.(*defs.SimpleSession).Username, false
  }
  return "", true
}