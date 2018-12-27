package main

import "testing"

func Testmax(t *testing.T) {
	if max(1, 2, 3, 4, 5) != 5 {
		t.Error("测试结果错误最大值寻找失败，max函数编写异常")
	} else {
		t.Log("测试结果成功，测试max函数结果正确")
	}
}

func TestMin(t *testing.T) {
	if min(1, 2, 3, 4, 5) != 1 {
		t.Error("测试结果失败，寻找最小值错误min函数编写异常")
	} else {
		t.Log("测试结果成功，寻找最小值min函数编程正确")
	}
}
