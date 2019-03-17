// 流控，连接数限制， bucket token算法

package main

import (
  "log"
)

type ConnLimiter struct {
  concurrentConn int   //并发连接数
  bucket chan int  //bucket channel数量
}

// buffer channel和普通channel
func NewConnLimiter(cc int) *ConnLimiter {
  return &ConnLimiter {
    concurrentConn: cc,
    bucket: make(chan int, cc),
  }
}

func (cl *ConnLimiter)  GetConn() bool {
  if len(cl.bucket) >= cl.concurrentConn {
    log.Printf("Reached the rate limitation.")
  }
  cl.bucket <- 1
  return true
}

func (cl *ConnLimiter) ReleaseConn(){
  c :=<- cl.bucket //释放操作
  log.Printf("New connection coming: %d", c)
}