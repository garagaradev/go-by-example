package main
import "fmt"

type Person struct {
  Name string
  Age int
}

func newPerson(name string) *Person {
  p := Person{Name: name}
  p.Age = 42
  return &p
}

func main(){
  fmt.Println(Person{"Bob",20})
  fmt.Println(Person{Name:"Alice",Age:30})
  fmt.Println(Person{Name:"Fred"})
  fmt.Println(&Person{Name:"Ann",Age:40})

  fmt.Println(newPerson("Jack"))

  s := Person{Name:"Sean",Age:50}
  fmt.Println(s.Name)

  sp := &s
  fmt.Println(sp.Age)

  sp.Age = 51
  fmt.Println(sp.Age)

  dog := struct {
    Name string
    isGood bool
  }{
    "Rex",
    true,
  }

  fmt.Println(dog)
}
