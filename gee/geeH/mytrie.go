package geeH

type node struct {
	pattern  string
	part     string
	children []*node
	isWide   bool
}

func (n *node) macthChild(path string) *node {
	for _, child := range n.children {
		if child.part == path || child.isWide {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(path string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == path || child.isWide {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
