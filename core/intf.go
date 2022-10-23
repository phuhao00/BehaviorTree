package core

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
)

type IBaseWrapper interface {
	_execute(tick *Tick) bt.Status
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) bt.Status
	_close(tick *Tick)
	_exit(tick *Tick)
}

type IBaseNode interface {
	IBaseWrapper

	Ctor()
	Initialize(params *config.BTNodeCfg)
	GetCategory() string
	Execute(tick *Tick) bt.Status
	GetName() string
	GetTitle() string
	SetBaseNodeWorker(worker IBaseWorker)
	GetBaseNodeWorker() IBaseWorker
}
