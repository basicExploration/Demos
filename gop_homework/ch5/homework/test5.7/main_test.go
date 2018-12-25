package main

import "testing"

func TestOutLine(t *testing.T) {

	err := outline("https:coastroad.net")
	if err != nil {
		t.Error(err)
	}
}

//--- FAIL: TestOutLine (0.00s)
//    main_test.go:9: Get https:coastroad.net: http: no Host in request URL
//FAIL
//exit status 1
//FAIL	github.com/googege/gopl_homework/ch5/homework/test5.7	0.019s
