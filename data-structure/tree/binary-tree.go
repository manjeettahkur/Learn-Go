package tree

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func Insert(t *Tree, val int) *Tree {

	if t == nil {
		return &Tree{Value: val, Left: nil, Right: nil}
	}

	if val < t.Value {
		t.Left = Insert(t.Left, val)
	} else {
		t.Right = Insert(t.Right, val)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}

	s := ""

	if t.Left != nil {
		s += t.Left.String() + " "
	}

	s += fmt.Sprint(t.Value)

	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"

}

func (t *Tree) WalkInorder() {
	if t == nil {
		return
	}

	t.Left.WalkInorder()
	fmt.Println("Node value", t.Value)
	t.Right.WalkInorder()

}

func (t *Tree) MaxDepth() int {
	if t == nil {
		return 0
	}

	lDepth := t.Left.MaxDepth()
	rDepth := t.Right.MaxDepth()

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func (t *Tree) WalkPreOrder() {
	if t == nil {
		return
	}

	fmt.Println("Node Value:", t.Value)
	t.Left.WalkPreOrder()
	t.Right.WalkPreOrder()
}

func (t *Tree) WalkPostOrder() {
	if t == nil {
		return
	}
	t.Left.WalkPostOrder()
	t.Right.WalkPostOrder()
	fmt.Println("Node Value:", t.Value)
}

func (t *Tree) Search(val int) {
	if t == nil {
		return
	}

	if val == t.Value {
		fmt.Println("Element found:", t.Value)
	}

	if val < t.Value {
		t.Left.Search(val)
	}
	t.Right.Search(val)
}

func (t *Tree) Delete(val int) *Tree {
	if t == nil {
		return nil
	}

	if val < t.Value {
		if t.Left != nil {
			t.Left = t.Left.Delete(val)
		}
	} else if val > t.Value {
		if t.Right != nil {
			t.Right = t.Right.Delete(val)
		}
	} else {

		if t.Left != nil && t.Right != nil {
			min := t.Min(t.Right)
			t.Value = min.Value
			t.Right = t.Right.Delete(min.Value)

		} else {
			// if no child
			if t.Left == nil && t.Right == nil {
				t = nil
			} else if t.Left == nil {
				t = t.Right
			} else {
				t = t.Left
			}
		}

	}

	return t
}

func (t *Tree) Min(n *Tree) *Tree {
	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}

	return n
}
