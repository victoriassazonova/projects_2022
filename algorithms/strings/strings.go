package strings

import "errors"

func LCS(s1, s2 string) string {
	m1 := []rune(s1)
	m2 := []rune(s2)
	aLen := len(m1)
	bLen := len(m2)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if m1[i] == m2[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x -= 1
		} else if lengths[x][y] == lengths[x][y-1] {
			y -= 1
		} else {
			s = append(s, m1[x-1])
			x -= 1
			y -= 1
		}
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func Segmentation(s string, isWord func(w string) bool) ([]string, error) {
	if len(s) == 0 {
		return []string{}, nil
	}
	res := make([]bool, len(s)+1)
	n := len(s)
	res[n] = true

	for r := n; r > 0; r-- {
		if res[r] {
			for l := 0; l < r; l++ {
				res[l] = isWord(s[l:r])

			}
		}

	}
	if !res[0] {
		return []string{}, errors.New("wrong")
	}
	words := make([]string, 0)
	start := 0
	end := 0
	for t := 0; t < len(res); t++ {
		if res[t] && t > end {
			end = t
			words = append(words, s[start:end])
			start = end
		}
	}
	return words, nil
}
