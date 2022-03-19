//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

package matrix

import "github.com/cheekybits/genny/generic"

type ValueType generic.Type

type Matrix struct {
	Rows   ValueType
	Cols   ValueType
	matrix []ValueType
}

func New(n, m ValueType) *Matrix {
	return &Matrix{
		Rows:   n,
		Cols:   m,
		matrix: make([]ValueType, n*m),
	}
}

func (M *Matrix) Get(i, j ValueType) ValueType {
	if i < 0 || i >= M.Rows || j < 0 || j >= M.Cols {
		panic("index error")
	}
	idx := M.Cols*i + j
	return M.matrix[idx]
}

func (M *Matrix) Set(i, j ValueType, v ValueType) {
	if i < 0 || i >= M.Rows || j < 0 || j >= M.Cols {
		panic("index error")
	}
	idx := M.Cols*i + j
	M.matrix[idx] = v
}
