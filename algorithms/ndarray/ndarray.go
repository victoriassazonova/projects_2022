package ndarray

type NDArray struct {
	shape []int
}

func New(shape ...int) *NDArray {
	for i := 0; i < len(shape); i++ {
		if shape[i] < 0 {
			panic("out of range")
		}
	}
	return &NDArray{shape}
}

func (nda *NDArray) Idx(indicies ...int) int {
	idx := 0

	if len(indicies) != len(nda.shape) {
		panic("out of range")
	}

	for i := 0; i < len(indicies); i++ {
		if indicies[i] < 0 || indicies[i] >= nda.shape[i] {
			panic("wrong value")
		}
		if i < len(indicies)-1 {
			c := indicies[i]
			for s := i + 1; s < len(nda.shape); s++ {
				c = c * nda.shape[s]
			}
			idx += c
		}
	}
	idx += indicies[len(indicies)-1]
	return idx
}
