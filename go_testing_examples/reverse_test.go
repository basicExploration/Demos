package reverse

import "testing"

type between struct {
	in string
	out string
}


func TestReverse(t *testing.T) {

	c := []*between{
		&between{
			"bta",
			"atb",
		},
		&between{
			"ttyu",
			"uytd",
		},
	}

	for k,v := range c {
		s := Reverse(v.in)
		if s != v.out {
			t.Log("这个选项测试错误:,错误值是:,期望值是:",k,s,v.out)
			t.Fail()
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Reverse("bta")
	}
}

