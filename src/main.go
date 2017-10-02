package main










func main() {
  printStartScreen()
  commandLoop()

  /*
  fnames := []string{"src/main.go", "src/filemanager.go", "src/folding.go"}
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

    dbpar := make(map[string]BitVect)
    db := NameDB{dbpar}
    for i:=0; i<len(fs1); i++ {
      db.addFile(fnames[i], fs1[i])
    }

    fmt.Println(db.extractArray())

  }
  */


}
