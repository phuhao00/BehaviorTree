package loader

import (
	"github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/actions"
	"github.com/phuhao00/BehaviorTree/composites"
	"github.com/phuhao00/BehaviorTree/config"
	"github.com/phuhao00/BehaviorTree/core"
	decorators "github.com/phuhao00/BehaviorTree/decorator"
)

func createBaseStructMaps() *BehaviorTree.RegisterStructMaps {
	st := BehaviorTree.NewRegisterStructMaps()
	//actions
	st.Register("Error", &actions.Error{})
	st.Register("Failure", &actions.Failure{})
	st.Register("Runner", &actions.Runner{})
	st.Register("Succeeder", &actions.Succeeder{})
	st.Register("Wait", &actions.Wait{})
	st.Register("Log", &actions.Log{})
	//composites
	st.Register("MemPriority", &composites.MemPriority{})
	st.Register("MemSequence", &composites.MemSequence{})
	st.Register("Priority", &composites.Priority{})
	st.Register("Sequence", &composites.Sequence{})

	//decorators
	st.Register("Inverter", &decorators.Inverter{})
	st.Register("Limiter", &decorators.Limiter{})
	st.Register("MaxTime", &decorators.MaxTime{})
	st.Register("Repeater", &decorators.Repeater{})
	st.Register("RepeatUntilFailure", &decorators.RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &decorators.RepeatUntilSuccess{})
	return st
}

func CreateBevTreeFromConfig(config *config.BTTreeCfg, extMap *BehaviorTree.RegisterStructMaps) *core.BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := core.NewBeTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}
