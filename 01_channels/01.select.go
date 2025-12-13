package main
import "fmt"
import "time"


func main() {
    ch1 := make (chan string)
    ch2 := make( chan string)
    go func(){
      time.Sleep(time.Second * 2)
       ch1 <- "ch1"
      
    }()
    go func (){
       time.Sleep(time.Second * 1)
       ch2 <- "ch2"
    }()
    select{
      case msg1 := <-ch1 :{
        fmt.Println(msg1)
        
      }
      case msg2 := <-ch2 :{
        fmt.Println(msg2)
        
      }
      
    }
     
}
