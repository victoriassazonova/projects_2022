package set

type Element interface {
	// Less returns true if the element is less than the other
	Less(other Element) bool
	// Equal returns true if the element is equivalent to the other
	Equal(other Element) bool
}

type Iterator struct {
	// contains filtered or unexported fields
}

func (it *Iterator) Next() bool {}

func (it *Iterator) Prev() bool {}

func (it *Iterator) Value() Element {}
