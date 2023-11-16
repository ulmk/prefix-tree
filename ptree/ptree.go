package ptree

type Node struct {
	Len      int
	Complete bool
	SubNodes map[byte]*Node
}

func New() *Node {
	return &Node{
		SubNodes: make(map[byte]*Node),
	}
}

func (n *Node) Clear() {
	if n.Len > 0 {
		for _, sub := range n.SubNodes {
			sub.Clear()
		}
	}
	n.Len = 0
	n.Complete = false
}

func (n *Node) Merge(mn *Node) {
	n.Len += mn.Len
	n.Complete = n.Complete || mn.Complete

	for c, sub := range mn.SubNodes {
		if sub.Len == 0 {
			continue
		}
		if nSub, ok := n.SubNodes[c]; ok {
			nSub.Merge(sub)
		} else {
			n.SubNodes[c] = sub
		}
	}
}

func (n *Node) Insert(w string) {
	n.Len++

	if len(w) == 0 {
		n.Complete = true
		return
	}

	c := w[0]

	sub, ok := n.SubNodes[c]
	if !ok {
		sub = New()
		n.SubNodes[c] = sub
	}
	sub.Insert(w[1:])
}

func (n *Node) Walk(f func(node *Node)) {
	f(n)
	for _, sub := range n.SubNodes {
		sub.Walk(f)
	}
}
