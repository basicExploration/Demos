package method
//type相当于构建了一个模型，到调用函数的时候需要初始化
//引入包
import "fmt"

//制定常量

const (
	PI  = 3.1415926
)

type Human struct {
	name string
	age int
	phoneNumber float64

}
 func (hu Human) run(value int) (int, float64) {
	 return value + hu.age, PI
 }

 func (hu *Human) Run(value int) string{
	 return "nice"
 }


func init() {
	hu := Human{"dd", 17, 554.333}
	//不初始化是没有值的
	a:= hu.Run(12)//这里虽然并没有使用&hu但是系统会自动取&
	fmt.Print(a)
}
