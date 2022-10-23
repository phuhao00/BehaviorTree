package example

import (
	"github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
	"github.com/phuhao00/BehaviorTree/loader"
	"path/filepath"
)

type Npc struct {
	*core.BehaviorTree
	*core.Blackboard
	TickCount int64
}

func (npc *Npc) loadBehaviorTree(treeFile string) {
	treeCfg, ok := config.LoadTreeCfg(filepath.FromSlash(treeFile))
	if !ok {
		return
	}
	maps := BehaviorTree.NewRegisterStructMaps()
	maps.Register("MoveToTarget", new(MoveToTarget))
	npc.BehaviorTree = loader.CreateBevTreeFromConfig(treeCfg, maps)
	npc.Blackboard = core.NewBlackboard()
	npc.Blackboard.SetTree("owner", npc, npc.BehaviorTree.GetID())
}

func (npc *Npc) update() {
	if npc.BehaviorTree != nil {
		npc.TickCount++
		npc.BehaviorTree.Tick(npc.TickCount, npc.Blackboard)
	}
}

func (npc *Npc) moveToTarget() BehaviorTree.Status {
	return BehaviorTree.RUNNING
}
