package dbops

import (
  "log"
  "time"
  "database/sql"
  "xcvideo/api/defs"
  "xcvideo/api/utils"
)

// run go test or go test -v (打印详细测试信息)
func AddUserCredential(loginName string, pwd string) error {
  // Prepare为预编译
  stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) VALUES (?,?)")
  if err != nil {
    return err
  }
  _, err = stmtIns.Exec(loginName, pwd)
  if err != nil {
    return err
  }
  defer stmtIns.Close()
  return nil
}

func GetUserCredential(loginName string) (string, error) {
  stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
  if err != nil {
    log.Printf("%s", err)
    return "", err
  }
  var pwd string
  err = stmtOut.QueryRow(loginName).Scan(&pwd)
  if err != nil && err != sql.ErrNoRows {
    return "", err
  }
  defer stmtOut.Close()

  return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
  stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
  if err != nil {
    log.Printf("DeleteUser error: %s", err)
    return err
  }
  _, err =  stmtDel.Exec(loginName, pwd)

  if err != nil {
    return err
  }
  defer stmtDel.Close()
  return nil
}

// // video
func AddNewVideo(aid int, name string) (*defs.VideoInfos, error) {
   // create uuid
  uuid, err := utils.NewUUID()
  if err != nil {
    return nil, err
  }
  t := time.Now()
  ctime := t.Format("Jan 02 2006, 15:04:05") //该时间为go中的时间原点,不可改变,格式为: M D y, HH:MM:SS
  // 如果字符串中有换行，则使用``符号
  stmtIns, err := dbConn.Prepare("INSERT INTO video_infos(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?, ?)")
  if err != nil {
    return nil, err
  }
  _, err = stmtIns.Exec(uuid, aid, name, ctime)
  if err != nil {
    return nil, err
  }
  res := &defs.VideoInfos{Id: uuid, AuthorId: aid, Name: name, DisplayCtime: ctime}
  defer stmtIns.Close()
  return res, nil
}

func AddNewComments(vid string, aid int, content string) error {
  uuid, err := utils.NewUUID()
  if err != nil {
    return err
  }
  stmtIns, err := dbConn.Prepare("INSERT INTO comments(id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
  if err != nil {
    return  err
  }
  _, err = stmtIns.Exec(uuid, vid, aid, content)
  if err != nil {
    return  err
  }
  defer stmtIns.Close()
  return nil
}

// 返回一个slice,[]*defs.Comments
func ListComments(vid string, from, to int) ([]*defs.Comments, error) {
  //按时间区间查询时，起始时间作为开区间，结束时间作为闭区间，为了解决在插入当时如果有查询操作时不会漏掉当前时间创建的记录
  stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.login_name, comments.content FROM comments
      INNER JOIN users on comments.author_id = users.id
      WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

  var res []*defs.Comments

  rows, err := stmtOut.Query(vid, from, to)
  if err != nil {
    return res, err
  }

  for rows.Next() {
    var id, name, content string
    if err := rows.Scan(&id, &name, &content); err != nil {
      return res, err
    }
    c := &defs.Comments{Id: id, VideoId: vid, Author: name, Content: content}
    res = append(res, c)
  }
  defer stmtOut.Close()
  return res, nil
}