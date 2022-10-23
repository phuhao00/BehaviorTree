package actions

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Succeeder struct {
	core.Action
}

func (this *Succeeder) OnTick(tick *core.Tick) bt.Status {
	return bt.SUCCESS
}
