package intset

import (
	"fmt"
	"testing"
)

var intS = new(IntSet)

func init() {
	intS = &IntSet{
		[]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
}

func TestHas(t *testing.T) {
	var result uint64
	intS.Add(1)
	fmt.Println(intS.words)
	for _, v := range intS.words {
		result += v
	}
	if result != 57 {
		t.Error("验证失败，结果不是66")
	}

}

func TestAdd(t *testing.T) {
	if !intS.Has(0) {
		t.Error("验证失败，没有0")
	}
}

func TestUnionWith(t *testing.T) {
	intS.UnionWith(intS)
}
