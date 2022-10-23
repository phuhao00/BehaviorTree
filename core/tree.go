package core

import (
	bt "github.com/phuhao00/BehaviorTree"
	"github.com/phuhao00/BehaviorTree/config"
)

type BehaviorTree struct {
	id string

	title string

	description string

	properties map[string]interface{}

	root IBaseNode

	debug interface{}

	dumpInfo *config.BTTreeCfg
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

func (this *BehaviorTree) Initialize() {
	this.id = bt.CreateUUID()
	this.title = "The behavior tree"
	this.description = "Default description"
	this.properties = make(map[string]interface{})
	this.root = nil
	this.debug = nil
}

func (this *BehaviorTree) GetID() string {
	return this.id
}

func (this *BehaviorTree) GetTitile() string {
	return this.title
}

func (this *BehaviorTree) SetDebug(debug interface{}) {
	this.debug = debug
}

func (this *BehaviorTree) GetRoot() IBaseNode {
	return this.root
}

func (this *BehaviorTree) Load(data *config.BTTreeCfg, maps *bt.RegisterStructMaps, extMaps *bt.RegisterStructMaps) {
	this.title = data.Title             //|| this.title;
	this.description = data.Description // || this.description;
	this.properties = data.Properties   // || this.properties;
	this.dumpInfo = data
	nodes := make(map[string]IBaseNode)

	// Create the node list (without connection between them)

	for id, s := range data.Nodes {
		spec := &s
		var node IBaseNode

		if spec.Category == "tree" {
			node = new(SubTree)
		} else {
			if extMaps != nil && extMaps.CheckElem(spec.Name) {
				// Look for the name in custom nodes
				if tnode, err := extMaps.New(spec.Name); err == nil {
					node = tnode.(IBaseNode)
				}
			} else {
				if tnode, err2 := maps.New(spec.Name); err2 == nil {
					node = tnode.(IBaseNode)
				} else {
					//fmt.Println("new ", spec.Name, " err:", err2)
				}
			}
		}

		if node == nil {
			// Invalid node name
			panic("BehaviorTree.load: Invalid node name:" + spec.Name + ",title:" + spec.Title)

		}

		node.Ctor()
		node.Initialize(spec)
		node.SetBaseNodeWorker(node.(IBaseWorker))
		nodes[id] = node
	}

	// Connect the nodes
	for id, spec := range data.Nodes {
		node := nodes[id]

		if node.GetCategory() == bt.COMPOSITE && spec.Children != nil {
			for i := 0; i < len(spec.Children); i++ {
				var cid = spec.Children[i]
				comp := node.(IComposite)
				comp.AddChild(nodes[cid])
			}
		} else if node.GetCategory() == bt.DECORATOR && len(spec.Child) > 0 {
			dec := node.(IDecorator)
			dec.SetChild(nodes[spec.Child])
		}
	}

	this.root = nodes[data.Root]
}

/**
 * This method dump the current BT into a data structure.
 *
 * Note: This method does not record the current node parameters. Thus,
 * it may not be compatible with load for now.
 *
 * @method dump
 * @return {Object} A data object representing this tree.
**/
func (this *BehaviorTree) dump() *config.BTTreeCfg {
	return this.dumpInfo
}

func (this *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) bt.Status {
	if blackboard == nil {
		panic("The blackboard parameter is obligatory and must be an instance of b3.Blackboard")
	}

	/* CREATE A TICK OBJECT */
	var tick = NewTick()
	tick.debug = this.debug
	tick.target = target
	tick.Blackboard = blackboard
	tick.tree = this

	/* TICK NODE */
	var state = this.root._execute(tick)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	var lastOpenNodes = blackboard._getTreeData(this.id).OpenNodes
	var currOpenNodes []IBaseNode
	currOpenNodes = append(currOpenNodes, tick._openNodes...)

	// does not close if it is still open in this tick
	var start = 0
	for i := 0; i < bt.MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i + 1
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	// close the nodes
	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	/* POPULATE BLACKBOARD */
	blackboard._getTreeData(this.id).OpenNodes = currOpenNodes
	blackboard.SetTree("nodeCount", tick._nodeCount, this.id)

	return state
}

func (this *BehaviorTree) Print() {

}
