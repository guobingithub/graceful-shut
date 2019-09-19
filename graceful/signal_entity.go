package graceful

import (
	"fmt"
	"github.com/guobingithub/graceful-shut/logger"
	"os"
)

type GracefulShutObj struct {
	Name string
	Func func()
}

type SignalServer struct {
	listenOn map[os.Signal]int
	shutObjs []*GracefulShutObj
}

func NewSignal() *SignalServer {
	ss := new(SignalServer)
	ss.listenOn = make(map[os.Signal]int)
	ss.shutObjs = make([]*GracefulShutObj, 0, 10)
	return ss
}

// 注册信号
func (ss *SignalServer) ListenOn(s os.Signal) {
	if _, ok := ss.listenOn[s]; !ok {
		ss.listenOn[s] = 1
	}
}

func (ss *SignalServer) RegisterGracefulShutObj(obj *GracefulShutObj) {
	ss.shutObjs = append(ss.shutObjs, obj)
}

func (ss *SignalServer) GracefulHandle(s os.Signal) {
	if _, ok := ss.listenOn[s]; ok {
		for _, obj := range ss.shutObjs {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(fmt.Sprintf("GracefulHandle, graceful shut (%s) error:(%v)", obj.Name, err))
				}
			}()
			obj.Func()
			logger.Info(fmt.Sprintf("GracefulHandle, graceful shut (%s) success", obj.Name))
		}
	}
}
