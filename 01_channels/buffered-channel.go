
package main
import "fmt"

func fib (n int , ch chan int){
    a:=0
    b :=1
    for i :=0;i<n;i++{
        ch <-a
        c := a+b
        a = b
        b = c
    }
    close(ch)
    
}
func main() {
    ch := make (chan int , 10)
    capacity :=cap(ch)
    go fib (capacity, ch)
    for val := range ch{
        fmt.Println(val)
    }
}
