package fulltext

import (
	"strings"
)

//
//type Index struct {
//	idx  map[string][]int
//	docs int
//}

//func New(docs []string) *Index {
//	idx := new(Index)
//	idx.docs = len(docs)
//	idx.idx = make(map[string][]int)
//	for i := 0; i < idx.docs; i++ {
//		words := strings.Fields(docs[i])
//		for j := 0; j < len(words); j++ {
//			if idx.idx[words[j]] == nil {
//				idx.idx[words[j]] = make([]int, idx.docs)
//			}
//			idx.idx[words[j]][i] = 1
//		}
//	}
//	return idx
//}
//
//func (idx *Index) Search(query string) []int {
//	wordsss := strings.Fields(query)
//	c := make([]int, idx.docs)
//	ans := make([]int, 0)
//	if len(query) == 0 {
//		return make([]int, 0)
//	}
//	for k := 0; k < idx.docs; k++ {
//		for i := 0; i < len(wordsss); i++ {
//
//			//for k := 0; k < idx.docs; k++ {
//			if idx.idx[wordsss[i]] != nil {
//				c[k] += idx.idx[wordsss[i]][k]
//			} else {
//				break
//			}
//		}
//	}
//	for k := 0; k < len(c); k++ {
//		if c[k] == len(wordsss) {
//			ans = append(ans, k)
//		}
//	}
//
//	return ans
//}
//
//type Index struct {
//	idx map[int]map[string]bool
//}
//
//func New(docs []string) *Index {
//	idx := new(Index)
//	idx.idx = make(map[int]map[string]bool)
//	for i := 0; i < len(docs); i++ {
//		idx.idx[i] = make(map[string]bool)
//		words := strings.Fields(docs[i])
//		for j := 0; j < len(words); j++ {
//			idx.idx[i][words[j]] = true
//		}
//	}
//	return idx
//}
//
//func (idx *Index) Search(query string) []int {
//	wordsss := strings.Fields(query)
//	c := make([]int, 0)
//	ans := make([]int, len(idx.idx))
//	if len(query) == 0 {
//		return make([]int, 0)
//	}
//	for k := range idx.idx {
//		c = append(c, k)
//		if len(idx.idx[k]) < len(wordsss) {
//			ans[k] = 1
//		}
//		for i := 0; i < len(wordsss) && ans[k] == 0; i++ {
//			if idx.idx[k][wordsss[i]] {
//				continue
//			} else {
//				c = c[:len(c)-1]
//				ans[k] = 1
//			}
//		}
//	}
//	return c
//}

type Index struct {
	idx map[string]map[int]bool
	l   int
}

func New(docs []string) *Index {
	idx := new(Index)
	idx.idx = make(map[string]map[int]bool)
	idx.l = len(docs)
	for i := 0; i < len(docs); i++ {
		words := strings.Fields(docs[i])
		for j := 0; j < len(words); j++ {
			if idx.idx[words[j]] == nil {
				idx.idx[words[j]] = make(map[int]bool)
			}
			idx.idx[words[j]][i] = true
		}
	}
	return idx
}

//func (idx *Index) Search(query string) []int {
//	wordsss := strings.Fields(query)
//	c := make([]int, 0)
//	ans := make([]int, idx.l)
//	if len(query) == 0 {
//		return make([]int, 0)
//	}
//	for i := 0; i < len(wordsss); i++ {
//		if idx.idx[wordsss[i]] == nil {
//			return make([]int, 0)
//		}
//		for k := 0; k < idx.l; k++ {
//			if idx.idx[wordsss[i]][k] {
//				continue
//			} else {
//				ans[k] = -1
//			}
//		}
//	}
//	for a := 0; a < len(ans); a++ {
//		if ans[a] > -1 {
//			c = append(c, a)
//		}
//	}
//	return c
//}
func (idx *Index) Search(query string) []int {
	wordsss := strings.Fields(query)
	c := make([]int, 0)
	ans := make(map[int]bool)
	for value := 0; value < idx.l; value++ {
		ans[value] = true
	}

	if len(query) == 0 {
		return make([]int, 0)
	}

	for i := 0; i < len(wordsss); i++ {
		if idx.idx[wordsss[i]] == nil {
			return make([]int, 0)
		}
		for k := range ans {
			if !idx.idx[wordsss[i]][k] {
				delete(ans, k)

			}
		}
		//fmt.Println(ans)
	}
	//fmt.Println(ans)
	for a, k := range ans {
		if k {
			c = append(c, a)
		}
	}
	return c
}
