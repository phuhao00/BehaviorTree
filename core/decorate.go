package core

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
)

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
	child IBaseNode
}

func (this *Decorator) Ctor() {

	this.category = bt.DECORATOR
}

func (this *Decorator) Initialize(params *config.BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
}

//GetChild
func (this *Decorator) GetChild() IBaseNode {
	return this.child
}

func (this *Decorator) SetChild(child IBaseNode) {
	this.child = child
}
