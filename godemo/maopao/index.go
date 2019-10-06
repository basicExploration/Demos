package maopao
// 冒泡排序并输出结果。
func numb(a []int64){
    for i := 0; i < len(a) -1; i ++ {
        c := false
        for j := 0; j < len(a)-1-i;j ++ {
           if a[j] >= a [j+1] {
            a[j+1],a[j] = a[j],a[j+1]
                   c = true
            }
        } 
        if !c  { /// 证明某次冒泡，全员不需要冒泡了，那么就证明已经排序OK了，就可以退出了。
            break
        }
    }
    for _,value := range a {
        fmt.Println(value)
    }

}
