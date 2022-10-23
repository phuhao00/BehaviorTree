package core

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
)

type BaseNode struct {
	IBaseWorker

	id string

	name string

	category string

	title string

	description string

	parameters map[string]interface{}

	properties map[string]interface{}
}

func (this *BaseNode) Ctor() {

}

func (this *BaseNode) SetName(name string) {
	this.name = name
}
func (this *BaseNode) SetTitle(name string) {
	this.name = name
}

func (this *BaseNode) SetBaseNodeWorker(worker IBaseWorker) {
	this.IBaseWorker = worker
}

func (this *BaseNode) GetBaseNodeWorker() IBaseWorker {
	return this.IBaseWorker
}

func (this *BaseNode) Initialize(params *config.BTNodeCfg) {
	//this.id = b3.CreateUUID()
	//this.title       = this.title || this.name
	this.description = ""
	this.parameters = make(map[string]interface{})
	this.properties = make(map[string]interface{})

	this.id = params.Id //|| node.id;
	this.name = params.Name
	this.title = params.Title             //|| node.title;
	this.description = params.Description // || node.description;
	this.properties = params.Properties   //|| node.properties;

}

func (this *BaseNode) GetCategory() string {
	return this.category
}

func (this *BaseNode) GetID() string {
	return this.id
}

func (this *BaseNode) GetName() string {
	return this.name
}
func (this *BaseNode) GetTitle() string {
	//fmt.Println("GetTitle ", this.title)
	return this.title
}

func (this *BaseNode) _execute(tick *Tick) bt.Status {
	//fmt.Println("_execute :", this.title)
	// ENTER
	this._enter(tick)

	// OPEN
	if !tick.Blackboard.GetBool("isOpen", tick.tree.id, this.id) {
		this._open(tick)
	}

	// TICK
	var status = this._tick(tick)

	// CLOSE
	if status != bt.RUNNING {
		this._close(tick)
	}

	// EXIT
	this._exit(tick)

	return status
}
func (this *BaseNode) Execute(tick *Tick) bt.Status {
	return this._execute(tick)
}

func (this *BaseNode) _enter(tick *Tick) {
	tick._enterNode(this)
	this.OnEnter(tick)
}

func (this *BaseNode) _open(tick *Tick) {
	//fmt.Println("_open :", this.title)
	tick._openNode(this)
	tick.Blackboard.Set("isOpen", true, tick.tree.id, this.id)
	this.OnOpen(tick)
}

func (this *BaseNode) _tick(tick *Tick) bt.Status {
	//fmt.Println("_tick :", this.title)
	tick._tickNode(this)
	return this.OnTick(tick)
}

func (this *BaseNode) _close(tick *Tick) {
	tick._closeNode(this)
	tick.Blackboard.Set("isOpen", false, tick.tree.id, this.id)
	this.OnClose(tick)
}

func (this *BaseNode) _exit(tick *Tick) {
	tick._exitNode(this)
	this.OnExit(tick)
}
