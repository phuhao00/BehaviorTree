package actions

import (
	"github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Error struct {
	core.Action
}

func (this *Error) OnTick(tick *core.Tick) BehaviorTree.Status {
	return BehaviorTree.ERROR
}
