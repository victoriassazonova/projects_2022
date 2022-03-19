package radix

func Sort(data []uint64) {
	digits := 8
	s := make([]uint64, len(data))
	for i := 0; i < digits; i++ {
		s = counting_sort(data, s, i)
		for q := 0; q < len(s); q++ {
			data[q] = s[q]
		}
	}
}

func counting_sort(abs []uint64, B []uint64, digit int) []uint64 {
	C := make([]uint64, 256)
	var d uint64
	var digit_of_Ai uint8
	for i := 0; i < len(abs); i++ {
		d = abs[i]
		d >>= 8 * digit
		digit_of_Ai = uint8(d)
		C[digit_of_Ai] = C[digit_of_Ai] + 1
	}
	for j := 1; j < 256; j++ {
		C[j] = C[j] + C[j-1]
	}
	for m := len(abs) - 1; m > -1; m-- {
		d = abs[m]
		d >>= 8 * digit
		digit_of_Ai = uint8(d)
		C[digit_of_Ai] = C[digit_of_Ai] - 1
		B[C[digit_of_Ai]] = abs[m]
	}
	C = nil
	return B
}
