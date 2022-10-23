package actions

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Runner struct {
	core.Action
}

func (this *Runner) OnTick(tick *core.Tick) bt.Status {
	return bt.RUNNING
}
