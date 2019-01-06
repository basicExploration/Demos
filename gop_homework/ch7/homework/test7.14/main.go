//练习 7.14： 定义一个新的满足Expr接口的具体类型并且提供一个新的操作例如对它运算单元中的最小值的计算。
// 因为Parse函数不会创建这个新类型的实例，为了使用它你可能需要直接构造一个语法树（或者继承parser接口）。
package main


func main(){

}


type str struct {

}

func(s *str)Eval(env Env) float64 {
	return 1.0
}

func(s *str)Check(vars map[Var]bool) error{
return nil
}