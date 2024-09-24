package main
import "fmt"

func main(){

  messages := make(chan string, 2)

  go func (){
    messages <- "buffered"
    messages <- "channel"
    messages <- "wrong message"
  }() //goroutine 2


  fmt.Println(<-messages)
  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
