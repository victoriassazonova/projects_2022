//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

package maxqueue

import (
	"errors"
	"github.com/cheekybits/genny/generic"
)

type ValueType generic.Type

type MaxQueue struct {
	stack1 []ValueType
	stack2 []ValueType
	max1   []ValueType
	max2   []ValueType
}

func New() *MaxQueue {
	return &MaxQueue{
		stack1: []ValueType{},
		stack2: []ValueType{},
		max1:   []ValueType{},
		max2:   []ValueType{},
	}
}

func (Mq *MaxQueue) Pop() (ValueType, error) {
	if len(Mq.stack1) == 0 && len(Mq.stack2) == 0 {
		return 0, errors.New("empty")
	}
	if Mq.max1[0] == Mq.max2[0] {
		Mq.max1 = Mq.max1[1:]
		Mq.max2 = Mq.max2[1:]
	} else {
		Mq.max1 = Mq.max1[1:]
	}
	if len(Mq.stack2) == 0 {
		for p := len(Mq.stack1) - 1; p > -1; p = p - 1 {
			Mq.stack2 = append(Mq.stack2, Mq.stack1[p])
			Mq.stack1 = Mq.stack1[:p]
		}
	}
	first := 0
	if len(Mq.stack2) != 0 {
		first := Mq.stack2[len(Mq.stack2)-1]
		Mq.stack2 = Mq.stack2[:len(Mq.stack2)-1]
		return first, nil
	}
	return first, nil
}

func (Mq *MaxQueue) Push(value ValueType) {
	if len(Mq.stack1) == 0 {
		Mq.stack1 = append(Mq.stack1, value)
		Mq.max1 = append(Mq.max1, value)
		Mq.max2 = append(Mq.max2, value)
	} else {
		Mq.stack1 = append(Mq.stack1, value)
		Mq.max1 = append(Mq.max1, value)
		for p := len(Mq.max2) - 1; p > -1; p = p - 1 {
			if value > Mq.max2[p] {
				Mq.max2 = Mq.max2[:p]
			} else {
				break
			}
		}
		Mq.max2 = append(Mq.max2, value)
	}
}

func (Mq *MaxQueue) Max() (ValueType, error) {
	if len(Mq.stack1) == 0 && len(Mq.stack2) == 0 {
		return 0, errors.New("empty")
	}
	return Mq.max2[0], nil
}
