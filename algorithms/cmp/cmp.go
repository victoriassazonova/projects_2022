//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

package cmp

import "github.com/cheekybits/genny/generic"

type ValueType generic.Number

func Min(values ...ValueType) ValueType {
	if len(values) == 0 {
		panic("at least one argument required")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}
	return m
}

func Max(values ...ValueType) ValueType {
	if len(values) == 0 {
		panic("at least one argument required")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}
	return m
}
