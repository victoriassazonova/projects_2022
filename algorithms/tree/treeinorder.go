package tree

func (T *Tree) InOrder(visit func(node *Tree)) {
	if T == nil {
		return
	}
	current := T
	stack := make([]*Tree, 0)
	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)

			current = current.Left
		} else {
			current = stack[len(stack)-1]
			visit(current)
			stack = stack[:len(stack)-1]

			current = current.Right
		}
	}
}
