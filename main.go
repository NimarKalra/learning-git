package main

import "fmt"

func f(x int) {
	if x == 0 {
		return
	}
	fmt.Print("hi\n")
	f(x - 1)
}
func main() {

	f(5)

}
