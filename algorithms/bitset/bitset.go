package bitset

import (
	"errors"
	"math/bits"
)

type Bitset struct {
	set    []uint64
	length int
}

func New(size int) *Bitset {
	bset := new(Bitset)
	n := 1 + (size-1)/64
	bset.set = make([]uint64, n)
	bset.length = size
	return bset
}

func (b *Bitset) All() bool {
	if b.length > 0 {
		return b.Count() == b.length
	}
	return false
}

func (b *Bitset) Any() bool {
	if b.length > 0 {
		return b.Count() > 0
	}
	return false
}

//func (b *Bitset) Flip() {
//
//	for w, x := range b.set {
//		if x == 1 {
//			b.Set(w, false)
//		}
//		if x == 0 {
//			b.Set(w, true)
//		}
//	}
//}
//func (b *Bitset) Flip() {
//	for w := range b.set {
//		for i := 0; i < 8; i++ {
//			current := b.set[w]
//			if b.set[w>>i] != 0 {
//				current = current &^ (1 << i)
//			} else {
//				current = current | (1 << i)
//			}
//			if current != 0 {
//				b.set[w] = current
//			} else {
//				current = current | (0 << i)
//				b.set[w] = current
//			}
//
//		}
//	}
//}

//func (b *Bitset) Flip() {
//	for w := range b.set {
//		current := b.set[w]
//
//		b.set[w] = 255 - current
//
//	}
//}

//func (b *Bitset) Flip() {
//	for w := range b.set {
//		current := b.set[w]
//		if current == 0 {
//			current = 255
//			b.set[w] = current
//		} else {
//			for i := 0; i < 8; i++ {
//				current = current | (0 << i)
//				if current>>i == 0 {
//					current = current | (1 << i)
//				} else {
//					current = current &^ (1 << i)
//
//				}
//				b.set[w] = current
//			}
//		}
//
//		fmt.Println(b.set[w])
//
//	}
//}

func (b *Bitset) Flip() {
	for w := range b.set {
		current := b.set[w]
		if current == 0 {
			b.set[w] = 18446744073709551615
		} else if current == 18446744073709551615 {
			b.set[w] = 0
		} else {

			for i := 0; i < 64; i++ {

				if current&(1<<uint(i)) == 0 {

					current = current | (1 << i)

				} else {
					current = current &^ (1 << i)

				}
			}
			//fmt.Println(b.set[w])

			b.set[w] = current
		}
	}
}

func (b *Bitset) Reset() {
	for w := range b.set {
		b.set[w] = 0
	}

}

func (b *Bitset) Count() int {
	if b != nil && b.set != nil {
		cnt := 0
		for _, x := range b.set {
			cnt += bits.OnesCount64(x)
		}
		return int(cnt)
	}
	return 0
}

func (b *Bitset) Set(pos int, value bool) error {
	if pos >= b.length || pos < 0 {
		return errors.New("out of range")
	}
	index := pos / 64
	off := uint(pos % 64)
	current := b.set[index]
	if value {
		current = current | (1 << off)
	} else {
		current = current &^ (1 << off)
	}
	if current != 0 {
		b.set[index] = current
	} else {
		current = current | (0 << off)
		b.set[index] = current
	}
	return nil
}

func (b *Bitset) Test(pos int) (bool, error) {
	idx := pos / 64
	off := pos % 64
	if pos >= b.length || pos < 0 {
		return false, errors.New("out of range")
	}
	if ((b.set[idx] >> off) & 1) == 1 {
		return true, nil
	}
	return false, nil
}
