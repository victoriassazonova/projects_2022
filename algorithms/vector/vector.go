//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

package vector

import (
	"github.com/cheekybits/genny/generic"
	"math"
)

type ValueType generic.Type

type Vector struct {
	Len    int
	vector []ValueType
}

func New(cap ValueType) *Vector {
	return &Vector{
		Len:    0,
		vector: make([]ValueType, cap),
	}
}

func (a *Vector) Get(idx ValueType) ValueType {
	if idx >= a.Len {
		panic("")
	}
	return a.vector[idx]
}

func (a *Vector) Set(idx int, x ValueType) {
	if idx >= a.Len {
		panic("")
	}
	a.vector[idx] = x
}

func (a *Vector) Delete(idx ValueType) {
	if idx >= a.Len {
		panic("")
	}
	for i := idx; i < a.Len-1; i++ {
		a.vector[i] = a.vector[i+1]
	}
	//мб лишняя
	a.vector[a.Len-1] = 0
	a.Len -= 1
	if a.Len <= len(a.vector)/4*2 {
		temp := make([]ValueType, int(math.Ceil(float64(len(a.vector))/4*2)))
		for i := 0; i < a.Len; i++ {
			temp[i] = a.vector[i]
		}
		a.vector = temp
	}
}

func (a *Vector) Insert(idx int, x ValueType) {
	if a.Len == 0 && len(a.vector) == 0 {
		a.vector = make([]ValueType, 1)
		a.vector[idx] = x
		a.Len += 1
	} else {
		if a.Len == len(a.vector) {
			temp := make([]ValueType, int(math.Ceil(float64(len(a.vector))*4/3)))
			for i := 0; i < a.Len; i++ {
				temp[i] = a.vector[i]
			}
			a.vector = temp
		}
		for i := a.Len - 1; i >= idx && idx != a.Len; i-- {
			a.vector[i+1] = a.vector[i]
		}
		a.vector[idx] = x
		a.Len += 1
	}
}
func (a *Vector) Push(x ValueType) {
	a.Insert(a.Len, x)
}
func (a *Vector) Pop() ValueType {
	m := a.Get(a.Len - 1)
	a.Delete(a.Len - 1)
	return m
}
