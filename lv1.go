package main
import "fmt"
var ch1=make(chan bool)
var ch2=make(chan bool)
var ch3=make(chan bool)
func main(){
		go num()
		go num2()
		<-ch1
}
func num(){
	for i:=1;i<100;i=i+2 {
	ch2<-true
		fmt.Println(i)
	<-ch3
	}
}
func num2() {
	for i := 2; i <= 100; i = i + 2 {
		<-ch2
		fmt.Println(i)
		ch3<-true
	}
		ch1 <- true
}