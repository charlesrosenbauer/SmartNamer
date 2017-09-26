package main



import (
    "fmt"
  )






func main() {
  s0 := "Hello World!"
  s1 := "Help Word!"
  vect0 := representID(s0)
  vect0  = representPosition("main.go", 15, vect0)
  vect1 := representID(s1)
  vect1  = representPosition("main.go", 115, vect1)
  fmt.Println(vectToString(vectIntersection(vect0, vect1)))

  fs, err := loadSourceFiles([]string{"src/main.go", "src/filemanager.go", "src/folding.go"})
  if err != nil {
    fmt.Println(err)
  }else{
    for i:=0; i<len(fs); i++ {
      fmt.Println(fs[i])
    }
  }
}
