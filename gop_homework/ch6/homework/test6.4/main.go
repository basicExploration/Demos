//*练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
package main

type IntSet struct {
	Words []uint64
}

//Elems method.
func (s *IntSet) Elems() []uint64 {
	result := make([]uint64, 0)
	for _, v := range s.Words {
		result = append(result, v)
	}
	return result
}
