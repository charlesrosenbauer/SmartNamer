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

  fnames := []string{"src/main.go", "src/filemanager.go", "src/folding.go", "std.bz"}
  fs0, err0 := loadSourceFiles(fnames)
  if err0 != nil {

    fmt.Println(err0)

  }else{

    var fs1 [][]StringPos

    for i:=0; i<len(fs0); i++ {
      fsTemp, err1 := getIds(fs0[i], fnames[i])

      if err1 != nil {
        fmt.Println(err1)
      }else{
        fs1 = append(fs1, fsTemp)
      }
    }

    for i:=0; i<len(fs1); i++ {
      fmt.Println(fs1[i])
    }

  }


}
