package taskrunner

import (
  "log"
)

type Runner struct {
  Controller controlChan
  Error controlChan
  Data dataChan
  dataSize int
  longLived bool
  Dispatcher fn
  Executor fn
}

func NewRunner(size int, longlived bool, d fn, e fn) {
  return &Runner {
    Controller: make(chan string, 1), //非阻塞的带buffer的channel
    Error: make(chan string, 1),
    Data: make(chan interface{}, size),
    longLived: longlived,
    dataSize: size,
    Dispatcher: d,
    Executor: e,
  }
}