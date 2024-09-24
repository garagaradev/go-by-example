package main
import "fmt"


type Rect struct {
  Width, Height int
}
//Pointer Receiver
func (r *Rect) area() int {
  return r.Width * r.Height
}
//Value Receiver
func (r Rect) perim() int {
  return 2 * (r.Width + r.Height)
}

func main(){
  r := Rect{Width:10,Height:5}

  fmt.Println("area: ",r.area())
  fmt.Println("perim: ",r.perim())

}
