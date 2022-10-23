package composites

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/core"
)

type Sequence struct {
	core.Composite
}

func (this *Sequence) OnTick(tick *core.Tick) bt.Status {
	//fmt.Println("tick Sequence :", this.GetTitle())
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != bt.SUCCESS {
			return status
		}
	}
	return bt.SUCCESS
}
