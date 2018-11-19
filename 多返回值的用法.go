// go的多返回值，可以直接将另一个函数的多返回值直接利用。
package main

func main() {

}


func one()(string,string,string){
	return "!","1","1"
}

func two()(string,string,string){
	return one()
}
