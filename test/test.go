package main
import (
	"fmt"
	"time"
)
 
func run(c chan int)  {
	data := <- c
	fmt.Println(data)
	c <- c
}
 
func run2(c chan int)  {
	data := <- c
	fmt.Println(data)
}

func main() {
	c := make(chan int)
	go run(c)
	c<-15
	go run2(c)
	time.Sleep(100 * time.Millisecond)
	//c<-18
	//select{}
}
