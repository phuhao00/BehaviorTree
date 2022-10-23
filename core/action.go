package core

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
)

type IAction interface {
	IBaseNode
}

type Action struct {
	BaseNode
	BaseWorker
}

func (this *Action) Ctor() {
	this.category = bt.ACTION
}
func (this *Action) Initialize(params *config.BTNodeCfg) {

	//this.id = b3.CreateUUID()
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.parameters = make(map[string]interface{})
	this.properties = make(map[string]interface{})
}
