package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 – x1
  b := y2 – y1
  return math.Sqrt(a*a + b*b)
}

var x="Hello World !!!!!"

var y [4]int

func main(){
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }

  var total int
  for i := 1; i <= 10; i++ {
    total +=i
    fmt.Println(total)
  }
  myX:=6
  pointerFunction(&myX)
  fmt.Println(myX)

  r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())
}

func pointerFunction (myX *int) {
  *myX=1234445
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
