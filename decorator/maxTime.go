package decorators

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
	"time"
)

type MaxTime struct {
	core.Decorator
	maxTime int64
}

func (this *MaxTime) Initialize(setting *config.BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.maxTime = setting.GetPropertyAsInt64("maxTime")
	if this.maxTime < 1 {
		panic("maxTime parameter in Limiter decorator is an obligatory parameter")
	}
}

func (this *MaxTime) OnOpen(tick *core.Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
}

func (this *MaxTime) OnTick(tick *core.Tick) bt.Status {
	if this.GetChild() == nil {
		return bt.ERROR
	}
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime int64 = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())
	var status = this.GetChild().Execute(tick)
	if currTime-startTime > this.maxTime {
		return bt.FAILURE
	}

	return status
}
