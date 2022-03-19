package levenshtein

type Levenshtein struct {
	str1   []rune
	str2   []rune
	matrix [][]int
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func New(src, dst string) *Levenshtein {
	lv := new(Levenshtein)
	lv.str1 = []rune(src)
	lv.str2 = []rune(dst)
	lv.matrix = make([][]int, len(lv.str1)+1)

	for i := 0; i < len(lv.str1)+1; i++ {
		lv.matrix[i] = make([]int, len(lv.str2)+1)
		lv.matrix[i][0] = i * 1
	}
	for j := 1; j < len(lv.str2)+1; j++ {
		lv.matrix[0][j] = j * 1
	}
	for i := 1; i < len(lv.str1)+1; i++ {
		for j := 1; j < len(lv.str2)+1; j++ {
			del := lv.matrix[i-1][j] + 1
			match := lv.matrix[i-1][j-1]
			if lv.str1[i-1] != lv.str2[j-1] {
				match += 1
			}
			ins := lv.matrix[i][j-1] + 1
			lv.matrix[i][j] = min(del, min(match,
				ins))
		}
	}
	return lv
}

func (ls *Levenshtein) Distance() int {
	return ls.matrix[len(ls.matrix)-1][len(ls.matrix[0])-1]
}

func (ls *Levenshtein) Transcript() string {
	return back(len(ls.matrix)-1, len(ls.matrix[0])-1, ls.matrix)
}

func back(i int, j int, matrix [][]int) string {
	if i > 0 && matrix[i-1][j]+1 == matrix[i][j] {
		return back(i-1, j, matrix) + "D"
	}
	if j > 0 && matrix[i][j-1]+1 == matrix[i][j] {
		return back(i, j-1, matrix) + "I"
	}
	if i > 0 && j > 0 && matrix[i-1][j-1]+1 == matrix[i][j] {
		return back(i-1, j-1, matrix) + "R"
	}
	if i > 0 && j > 0 && matrix[i-1][j-1] == matrix[i][j] {
		return back(i-1, j-1, matrix) + "M"
	}
	var str1 string
	return str1
}
