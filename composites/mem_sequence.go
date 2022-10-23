package composites

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type MemSequence struct {
	core.Composite
}

func (this *MemSequence) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *MemSequence) OnTick(tick *core.Tick) bt.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), this.GetID())
	for i := child; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)

		if status != bt.SUCCESS {
			if status == bt.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), this.GetID())
			}

			return status
		}
	}
	return bt.SUCCESS
}
