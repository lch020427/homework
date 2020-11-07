package main
import (
	"fmt"
	"os"
)
func main() {
	file, _ := os.Create("proverb.txt")
	file1,_:=os.Open("proverb.txt")
	defer file.Close()
	str := "don't communicate by sharing memory share memory by communicating"
	file.WriteString(str)
	var buf [128]byte
	n,_:=file1.Read(buf[:] )
	fmt.Println(string(buf[0:n]))
}