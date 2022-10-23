package core

type Memory struct {
	_memory map[string]interface{}
}

func NewMemory() *Memory {
	return &Memory{make(map[string]interface{})}
}

func (this *Memory) Get(key string) interface{} {
	return this._memory[key]
}
func (this *Memory) Set(key string, val interface{}) {
	this._memory[key] = val
}
func (this *Memory) Remove(key string) {
	delete(this._memory, key)
}

type TreeMemory struct {
	*Memory
	_treeData   *TreeData
	_nodeMemory map[string]*Memory
}

func NewTreeMemory() *TreeMemory {
	return &TreeMemory{NewMemory(), NewTreeData(), make(map[string]*Memory)}
}
