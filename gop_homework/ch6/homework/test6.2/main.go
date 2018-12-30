//练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
package main

type IntSet struct {
	words []uint64
}

func (i *IntSet) AddAll(sep ...int) {
	for _, v := range sep {
		i.Add(v)
	}
}
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
