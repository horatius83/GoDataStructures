package tree 

type Element struct {
	Value interface { }
}

type node struct {
	data Element
	left *Node
	right *Node
}

type Comparison func(x, y Element) bool

type Binary struct {
	root Node
	compare Comparison
}

func (tree Binary) Add(e Element) {
	if tree.root != nil {
		tree.Add(tree.root, e)
	} else {
		tree.root = Node { data = e }
	}
}

func (tree Binary) Add(current, additional Element) {
	if current != nil {
		var result = tree.compare(current, additional)
		if result == false {
			if current.left != nil {
				Add(current.left, additional)
			} else {
				current.left = new Node { data = additional }
			}
		} else {
			if current.right != nil {
				Add(current.right, additional)
			} else {
				current.right = new Node { data = additional }
			}
		}
	}
}
