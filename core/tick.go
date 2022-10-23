package core

type Tick struct {
	tree *BehaviorTree

	debug interface{}

	target interface{}

	Blackboard *Blackboard

	_openNodes []IBaseNode

	_openSubtreeNodes []*SubTree

	_nodeCount int
}

func NewTick() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

func (this *Tick) Initialize() {
	// set by BehaviorTree
	this.tree = nil
	this.debug = nil
	this.target = nil
	this.Blackboard = nil

	// updated during the tick signal
	this._openNodes = nil
	this._openSubtreeNodes = nil
	this._nodeCount = 0
}

func (this *Tick) GetTree() *BehaviorTree {
	return this.tree
}

func (this *Tick) _enterNode(node IBaseNode) {
	this._nodeCount++
	this._openNodes = append(this._openNodes, node)

	// TODO: call debug here
}

func (this *Tick) _openNode(node *BaseNode) {
	// TODO: call debug here
}

/**
 * Callback when ticking a node (called by BaseNode).
 * @method _tickNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _tickNode(node *BaseNode) {
	// TODO: call debug here
	//fmt.Println("Tick _tickNode :", this.debug, " id:", node.GetID(), node.GetTitle())
}

/**
 * Callback when closing a node (called by BaseNode).
 * @method _closeNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _closeNode(node *BaseNode) {
	// TODO: call debug here

	ulen := len(this._openNodes)
	if ulen > 0 {
		this._openNodes = this._openNodes[:ulen-1]
	}

}

func (this *Tick) pushSubtreeNode(node *SubTree) {
	this._openSubtreeNodes = append(this._openSubtreeNodes, node)
}
func (this *Tick) popSubtreeNode() {
	ulen := len(this._openSubtreeNodes)
	if ulen > 0 {
		this._openSubtreeNodes = this._openSubtreeNodes[:ulen-1]
	}
}

/**
 * return top subtree node.
 * return nil when it is runing at major tree
 *
**/
func (this *Tick) GetLastSubTree() *SubTree {
	ulen := len(this._openSubtreeNodes)
	if ulen > 0 {
		return this._openSubtreeNodes[ulen-1]
	}
	return nil
}

/**
 * Callback when exiting a node (called by BaseNode).
 * @method _exitNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _exitNode(node *BaseNode) {
	// TODO: call debug here
}

func (this *Tick) GetTarget() interface{} {
	return this.target
}