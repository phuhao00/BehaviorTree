package composites

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Priority struct {
	core.Composite
}

func (this *Priority) OnTick(tick *core.Tick) bt.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != bt.FAILURE {
			return status
		}
	}
	return bt.FAILURE
}
