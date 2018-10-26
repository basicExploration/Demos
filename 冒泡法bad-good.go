func maoSliceGood(sli []int) time.Duration {
	start := time.Now()
	for i := 1; i < len(sli); i++ {
		for j := 1; j < len(sli)-i-1; j++ {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}
	// fmt.Println("数据是:", sli)
	end := time.Now()
	return end.Sub(start)
}
func maoSliceBad(sli []int) time.Duration {
	start := time.Now()
	for i := 1; i < len(sli); i++ {
		for j := 1; j < len(sli); j++ {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}
	// fmt.Println("数据是:", sli)
	end := time.Now()
	return end.Sub(start)
}
func ceshi(n int, sli1 []int, sli2 []int) {
	var (
		a int
		b int
	)
	j := 0

	for i := 0; i < n; i++ {
		a = int(maoSliceBad(sli1))
		b = int(maoSliceGood(sli2))
		if a-b > 0 {
			j++
		}
	 fmt.Println("好的坏的分别是，快了：", maoSliceGood(sli2), maoSliceBad(sli1), maoSliceBad(sli1)-maoSliceGood(sli2))
	}
	fmt.Printf("good比bad好的概率是:%d分之%d\n", n, j)

}

#---
好的坏的分别是，快了： 54.60225ms 57.56397ms 2.637157ms
好的坏的分别是，快了： 54.854517ms 58.02214ms 3.015214ms
好的坏的分别是，快了： 54.52057ms 57.654257ms 3.291049ms
好的坏的分别是，快了： 54.615181ms 57.506324ms 3.984839ms
好的坏的分别是，快了： 54.593795ms 58.074065ms 2.935677ms
好的坏的分别是，快了： 54.552618ms 57.526483ms 2.854921ms
好的坏的分别是，快了： 54.604215ms 57.686433ms 2.886505ms
好的坏的分别是，快了： 54.73695ms 57.693106ms 3.208127ms
好的坏的分别是，快了： 54.528703ms 57.66883ms 2.978056ms
好的坏的分别是，快了： 54.525377ms 63.522413ms 3.187294ms
good比bad好的概率是:10分之10
