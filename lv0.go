package main
import (
	"fmt"
)
var (
	myres = make(map[int]int, 20)
)
func factorial(n int) {
	ch:=make(chan int,20)
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch<-n
	myres[n] = res
	<-ch
}
func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}