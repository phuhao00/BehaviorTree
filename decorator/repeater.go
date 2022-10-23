package decorators

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
)

type Repeater struct {
	core.Decorator
	maxLoop int
}

func (this *Repeater) Initialize(setting *config.BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

func (this *Repeater) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("i", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *Repeater) OnTick(tick *core.Tick) bt.Status {
	//fmt.Println("tick ", this.GetTitle())
	if this.GetChild() == nil {
		return bt.ERROR
	}
	var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
	var status = bt.SUCCESS
	for this.maxLoop < 0 || i < this.maxLoop {
		status = this.GetChild().Execute(tick)
		if status == bt.SUCCESS || status == bt.FAILURE {
			i++
		} else {
			break
		}
	}
	tick.Blackboard.Set("i", i, tick.GetTree().GetID(), this.GetID())
	return status
}
