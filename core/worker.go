package core

import (
	"fmt"
	bt "github.com/phuhao00/BehaviorTree"
)

type IBaseWorker interface {
	OnEnter(tick *Tick)

	OnOpen(tick *Tick)

	OnTick(tick *Tick) bt.Status

	OnClose(tick *Tick)

	OnExit(tick *Tick)
}
type BaseWorker struct {
}

func (this *BaseWorker) OnEnter(tick *Tick) {

}

func (this *BaseWorker) OnOpen(tick *Tick) {

}

func (this *BaseWorker) OnTick(tick *Tick) bt.Status {
	fmt.Println("tick BaseWorker")
	return bt.ERROR
}

func (this *BaseWorker) OnClose(tick *Tick) {

}

func (this *BaseWorker) OnExit(tick *Tick) {

}
