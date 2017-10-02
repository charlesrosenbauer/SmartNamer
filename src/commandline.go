package main



import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "errors"
)










func commandLoop(){

  // System State
  var (
    fnames []string
    errs   []error
  )

  // Command Loop State
  reader := bufio.NewReader(os.Stdin)
  cont := true

  for cont {
    fmt.Print(">>> ")
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println(err)
    }
    command := strings.Split(strings.Trim(text, "\n"), " ")

    if len(command) > 0 {
      switch command[0] {
      case "quit" :
        cont = false
      case "load" : {
        for i := 1; i < len(command); i++ {
          fnames = append(fnames, command[i])
        }
      }
      case "show-file-list" : {
        for _, v := range fnames {
          fmt.Println(v)
        }
      }
      default :
          errs = append(errs, errors.New("Unknown Command"))
      }
    }

    if len(errs) > 0 {
      for _, v := range errs {
        fmt.Println(v)
      }
    }

  }
}
