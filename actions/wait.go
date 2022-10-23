package actions

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
	"time"
)

type Wait struct {
	core.Action
	endTime int64
}

func (this *Wait) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.endTime = setting.GetPropertyAsInt64("milliseconds")
}

func (this *Wait) OnOpen(tick *core.Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
}

func (this *Wait) OnTick(tick *core.Tick) bt.Status {
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())
	//fmt.Println("wait:",this.GetTitle(),tick.GetLastSubTree(),"=>", currTime-startTime)
	if currTime-startTime > this.endTime {
		return bt.SUCCESS
	}

	return bt.RUNNING
}
