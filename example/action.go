package example

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
)

// MoveToTarget ...
type MoveToTarget struct {
	core.Action
	owner *Npc
}

// Initialize ...
func (action *MoveToTarget) Initialize(setting *config.BTNodeCfg) {
	action.Action.Initialize(setting)
}

// OnOpen ...
func (action *MoveToTarget) OnOpen(tick *core.Tick) {
	owner := tick.Blackboard.Get("owner", tick.GetTree().GetID(), "")
	action.owner = owner.(*Npc)
}

// OnTick ...
func (action *MoveToTarget) OnTick(tick *core.Tick) bt.Status {
	return action.owner.moveToTarget()
}
