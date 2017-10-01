package main



import (
  "fmt"
  "bufio"
  "os"
  "strings"
)










func commandLoop(){
  reader := bufio.NewReader(os.Stdin)
  cont := true
  for cont {
    fmt.Print(">>> ")
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println(err)
    }
    command := strings.Split(strings.Trim(text, "\n"), " ")

    fmt.Println(command)
    cont = (text != "quit\n")
  }
}
