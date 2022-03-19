package tree

func (T *Tree) IsSym() bool {
	if T == nil {
		return true
	}
	if T.Left == nil && T.Right == nil {
		return true
	}
	if T.Left == nil || T.Right == nil {
		return false
	}
	return check(T.Left, T.Right)
}

func check(left *Tree, right *Tree) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value && check(left.Left, right.Right) && check(left.Right, right.Left)
}
