package main

import "testing"

func TestTree(t *testing.T) {
	value := Tree()
	if value == "" {
		t.Error("测试错误")
	} else {
		t.Log("ok!")
	}
}

//go test -test.v // 使用-test.v 可以获取到 log的输出值。
//=== RUN   TestTree
//--- PASS: TestTree (0.00s)
//    main_test.go:10: ok!
//PASS
//ok  	github.com/googege/gopl_homework/ch7/homework/test7.3	0.006s
