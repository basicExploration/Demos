package maopao
// 冒泡排序并输出结果。
func numb(a []int64){
    for i := 0; i < len(a) -1; i ++ {
        for j := 0; j < len(a)-1-i;j ++ {
               if a[j] >= a [j+1] {
            a[j+1],a[j] = a[j],a[j+1]
        
        }
        } 
    }
    for _,value := range a {
        fmt.Println(value)
    }

}
