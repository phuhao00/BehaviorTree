package actions

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Failure struct {
	core.Action
}

func (f *Failure) OnTick(tick *core.Tick) bt.Status {
	return bt.FAILURE
}
