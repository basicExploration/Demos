//练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的交集，
// 我们再为IntSet实现另外的几个函数
// IntersectWith(交集：元素在A集合B集合均出现),
// DifferenceWith(差集：元素出现在A集合，未出现在B集合),
// SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)。
package main

type IntSet struct {
	words []uint64
}

// 交集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) []uint64 {
	result := make([]uint64, 0)
	var key int
	var l int
	if len(t.words) > len(s.words) {
		l = len(s.words)
	} else {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if t.words[i] != s.words[j] {
				key++
			}

		}
		if key != l {
			result = append(result,t.words[i])
		}

	}
	return result

}

func (s *IntSet) IntersectWith(t *IntSet) []uint64{
	result := make([]uint64, 0)
	var l int
	if len(t.words) > len(s.words) {
		l = len(s.words)
	} else {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		if t.words[i] == s.words[i] {
			result = append(result, t.words[i])
		}
	}
	return result

}
func (s *IntSet) SymmetricDifference(t *IntSet)[]uint64 {
	result := make([]uint64, 0)
	var key int
	var l int
	if len(t.words) > len(s.words) {
		l = len(s.words)
	} else {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if t.words[i] != s.words[j] {
				key++
			}

		}
		if key != l {
			result = append(result,t.words[i])
		}

	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if s.words[i] != t.words[j] {
				key++
			}

		}
		if key != l {
			result = append(result,s.words[i])
		}

	}
	return result
}
