package trie

type Tree struct {
	root  node
	count int
}

type node struct {
	value    string
	children [256]*node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Count() int {
	return t.count
}

func (t *Tree) Insert(s string) bool {
	return false
}

func (t *Tree) Find(s string) bool {
	return t.findNode(s) != nil
}

func (t *Tree) findNode(s string) *node {
	node := &t.root
	for node != nil {
		if s == "" {

		}
		if node.isLeaf() {
			if node.value == s {
				return node
			}
			return nil
		}
		node = node.traverse(s[0])
		s = s[1:]
	}
	return nil
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) traverse(firstByte byte) *node {
	return node.children[firstByte]
}
