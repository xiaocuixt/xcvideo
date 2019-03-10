package defs

//request
type UserCredential struct {
  Username string `json:"user_name"`
  Pwd string `json:"pwd"`
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