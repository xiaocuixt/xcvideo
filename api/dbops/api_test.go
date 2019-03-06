package dbops

import (
  "testing"
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate table)

func clearTables() {
  dbConn.Exec("truncate users")
  dbConn.Exec("truncate video_infos")
  dbConn.Exec("truncate comments")
  dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
  clearTables()
  m.Run()
  clearTables()
}

func TestUserWorkFlow(t *testing.T) {
  t.Run("Add", testAddUser)
  t.Run("Get", testGetUser)
  t.Run("Del", testDeleteUser)
  t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
  err := AddUserCredential("avenssi", "123")
  if err != nil {
    t.Errorf("Error of AddUser: %v", err)
  }
}

func testGetUser(t *testing.T) {
  pwd, err := GetUserCredential("avenssi")
  if pwd != "123" || err != nil {
    t.Errorf("Error of GetUser: %v", err)
  }
}

func testDeleteUser(t *testing.T) {
  err := DeleteUser("avenssi", "123")
  if err != nil {
    t.Errorf("Error of DeleteUser: %v", err)
  }
}

// 确认是否真正的被删除
func testRegetUser(t *testing.T) {
  pwd, err := GetUserCredential("avenssi")
  if err != nil {
    t.Errorf("Error of DeleteUser: %v", err)
  }
  if pwd != "" {
    t.Errorf("Delete user test failed")
  }
}