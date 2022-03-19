package lru

import "container/list"

type Cache struct {
	m   map[int]*list.Element
	cap int
	l   *list.List
}

type Pair struct {
	key, value int
}

func New(capacity int) *Cache {
	c := new(Cache)
	c.l = list.New()
	c.cap = capacity
	c.m = make(map[int]*list.Element)

	return c
}

func (cache *Cache) Get(key int) (int, bool) {
	node, ok := cache.m[key]
	if !ok {
		return 0, false
	}
	cache.l.MoveToFront(node)
	val := node.Value.(Pair).value
	return val, true
}

func (cache *Cache) Put(key int, value int) {
	elem, ok := cache.m[key]
	if !ok {
		if cache.l.Len() == cache.cap {
			el := cache.l.Back()
			delete(cache.m, el.Value.(Pair).key)
			cache.l.Remove(el)
		}
		cache.l.PushFront(Pair{
			key,
			value,
		})
		cache.m[key] = cache.l.Front()
	} else {
		if elem.Value.(Pair).value != value {
			elem.Value = Pair{
				key,
				value,
			}
		}
		cache.l.MoveToFront(elem)
		cache.m[key] = cache.l.Front()
	}
}
