
package main
import "fmt"
func sum (arr [] int, ch chan int ){
    res :=0
    for _, val := range arr{
        res = res + val
    }
    ch <-res
}

func main() {
    arr := [] int {1,2,4,3,1,2,6}
    ch := make (chan int )
    go sum (arr[0:len(arr)/2], ch)
    go sum (arr[len(arr)/2:], ch)
    x, y := <-ch, <-ch
    fmt.Println(x, y)
}
