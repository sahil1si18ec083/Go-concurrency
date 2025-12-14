// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
import "sync"
type Counter struct{
    v map[string]int
}
 func (c *Counter) inc(key string, mu *sync.Mutex, wg *sync.WaitGroup){
    defer wg.Done()
    mu.Lock()
    defer mu.Unlock()
    c.v[key]++
    
}

func main() {
  var mu sync.Mutex
  var wg sync.WaitGroup
  fmt.Println("Try programiz.pro")
  c:= Counter {v : make(map[string]int)}
  for i:=0;i<10000;i++{
      wg.Add(1)
     go  c.inc("some key", &mu, &wg)
  }
  wg.Wait()
  
  
  fmt.Println(c)
}
