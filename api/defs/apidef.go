package defs

//request
type UserCredential struct {
  Username string `json:"user_name"`
  Pwd string `json:"pwd"`
}


//通过json声明后，我们数据格式转为json时，会按照指定的格式:
// {
//   "success": "",
//   "session_id": ""
// }
//response
type SignedUp struct {
  Success bool `json:"success"`
  SessionId string `json:"session_id"`
}


// Data Model
type VideoInfos struct {
  Id string
  AuthorId int
  Name string
  DisplayCtime string
}

type Comments struct {
  Id string
  VideoId string
  Author string
  Content string
}

type SimpleSession struct {
  Username string  //login name
  TTL int64
}