package main
import "fmt"
//var i, j int = 1,3
func main() {
//	k := "Hello world"
	const Big = 1 << 100
	const Small = Big >> 99
		fmt.Printf("The type of Big is: %T\n",Big)
//	var c,python,java = true, false, "no!!!"
	//fmt.Println(i,c,j,python,java,k)
}
