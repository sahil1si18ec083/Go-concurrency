package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	x := <-ch
	fmt.Println(x)
}

// package main
// import "fmt"

// func ping (s string ,  pingChan chan string){
       
//   pingChan <-s
// }
// func pong (pingChan chan string, pongChan chan string){
       
//       val := <-pingChan
//      pongChan <-val
// }
// func main() {
  
//      pingChan := make(chan string)
// 	   pongChan := make(chan string)
	
// 	   go ping ("hello", pingChan)
// 	   go pong (pingChan, pongChan)
// 	   fmt.Println(<-pongChan)
	
// }
