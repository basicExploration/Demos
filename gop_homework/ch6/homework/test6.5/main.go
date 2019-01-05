//练习 6.5： 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，
// 可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
package main

const IS64  = 32 << (^uint(0) >> 63) // 判断是否是64位平台。// 这样就可以照顾是32和64位系统了。

type IntSet struct {
	words []uint// 这样就可以照顾是32和64位系统了。
}

func (s *IntSet) Add(x int) {
	word, bit := x/IS64, uint(x%IS64)// 这样就可以照顾是32和64位系统了。
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 没有后缀的 uint 就是指最少是32 它会根据系统选择 32或者是64  那么该除以几呢？使用32 << (^uint(0) >> 63) 即可判断。