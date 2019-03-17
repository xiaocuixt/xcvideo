package taskrunner

const (
  READY_TO_DISPATCH = "d"
  READY_TO_EXCUTE = "e"
  CLOSE = "c"
)
type controlChan chan string

type dataChan chan interface{}  //泛型，可以为int, string或者其它类型

type fn func(dc dataChan) error