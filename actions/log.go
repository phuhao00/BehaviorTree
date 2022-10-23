package actions

import (
	"fmt"
	BehaviorTree2 "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
)

type Log struct {
	core.Action
	info string
}

func (l *Log) Initialize(setting *config.BTNodeCfg) {
	l.Action.Initialize(setting)
	l.info = setting.GetPropertyAsString("info")
}

func (l *Log) OnTick(tick *core.Tick) BehaviorTree2.Status {
	fmt.Println("log:", l.info)
	return BehaviorTree2.SUCCESS
}
