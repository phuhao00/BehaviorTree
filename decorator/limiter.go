package decorators

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
)

type Limiter struct {
	core.Decorator
	maxLoop int
}

func (this *Limiter) Initialize(setting *config.BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

func (this *Limiter) OnTick(tick *core.Tick) bt.Status {
	if this.GetChild() == nil {
		return bt.ERROR
	}
	var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
	if i < this.maxLoop {
		var status = this.GetChild().Execute(tick)
		if status == bt.SUCCESS || status == bt.FAILURE {
			tick.Blackboard.Set("i", i+1, tick.GetTree().GetID(), this.GetID())
		}
		return status
	}

	return bt.FAILURE
}
