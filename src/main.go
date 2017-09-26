package main



import (
    "fmt"
  )






func main() {
  s := "Hello World!"
  vect := representID(s)
  vect  = representPosition("main.go", 15, vect)
  fmt.Println(vectToString(vect))
}
