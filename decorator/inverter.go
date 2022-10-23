package decorators

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Inverter struct {
	core.Decorator
}

func (this *Inverter) OnTick(tick *core.Tick) bt.Status {
	if this.GetChild() == nil {
		return bt.ERROR
	}

	var status = this.GetChild().Execute(tick)
	if status == bt.SUCCESS {
		status = bt.FAILURE
	} else if status == bt.FAILURE {
		status = bt.SUCCESS
	}

	return status
}
