package tree

type Element struct {
	Value interface{}
}

type node struct {
	data  *Element
	left  *node
	right *node
}

type Comparison func(x, y *Element) bool

type Visitor func(x *Element)

type Binary struct {
	root    *node
	compare Comparison
}

func (tree *Binary) Add(e *Element) {
	if tree.root != nil {
		tree.add(tree.root, e)
	} else {
		tree.root = &node{data: e}
	}
}

func (tree *Binary) add(current *node, additional *Element) {
	if current != nil {
		var result = (*tree).compare(current.data, additional)
		if result == false {
			if current.left != nil {
				tree.add(current.left, additional)
			} else {
				current.left = &node{data: additional}
			}
		} else {
			if current.right != nil {
				tree.add(current.right, additional)
			} else {
				current.right = &node{data: additional}
			}
		}
	}
}

func (tree *Binary) DepthFirst(f Visitor) {
	tree.depthFirst(tree.root, f)
}

func (tree *Binary) depthFirst(n *node, f Visitor) {
	if n != nil {
		f(n.data)
	}
	if n.left != nil {
		tree.depthFirst(n.left, f)
	}
	if n.right != nil {
		tree.depthFirst(n.right, f)
	}
}

func (tree *Binary) Count() int {
	count := 0
	counter := func(e *Element) { count += 1 }
	tree.DepthFirst(counter)
	return count
}
