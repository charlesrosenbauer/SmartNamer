package main



import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "errors"
)










func printHelpScreen() {
  printStartScreen()
  fmt.Println("\nCommands: ")
  fmt.Println("add-to-workspace [fs] : Add source file to current workspace. Does not actually load the file or perform safety checks.")
  fmt.Println("show-workspace        : Show all files currently in the workspace. They are not necessarily loaded, however.")
  fmt.Println("clear-workspace       : Remove all files from current workspace.")
  fmt.Println("load-files       [fs] : Loads contents of any files in the workspace into memory. Optionally can add/load additional files in the process.")
  fmt.Println("?                     : Show Help Screen (You Are Here)")
  fmt.Println("quit                  : Quit the program")

  fmt.Println("\n")
}










func printStartScreen() {
    fmt.Println("\nSmart Namer, created by Charles Rosenbauer")
    fmt.Println("https://github.com/charlesrosenbauer/SmartNamer for info.")
    fmt.Println("Enter \"?\" for help.\n")
}










func commandLoop() {

  // System State
  var (
    errs  []error
  )
  files := map[string]string{}

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

      case "add-to-workspace" : {
        for i := 1; i < len(command); i++ {
          files[command[i]] = ""
        }
      }

      case "show-workspace" : {
        for k, _ := range files {
          fmt.Println(k)
        }
      }

      case "clear-workspace" : {
        files = map[string]string{}
      }

      case "load-files" : {
        for i := 1; i < len(command); i++ {
          files[command[i]] = ""
        }
        var err error
        files, err = loadSourceFromMap(files)
        if err != nil {
          errs = append(errs, err)
        }
      }

      case "show-file-texts" : {
        if len(files) == 0 {
          fmt.Println("Nothing loaded yet.")
        }
        for _, v := range files {
          fmt.Println(v)
        }
      }

      case "?" :
        printHelpScreen()

      default :
          errs = append(errs, errors.New("Unknown Command"))
      }
    }

    if len(errs) > 0 {
      for _, v := range errs {
        fmt.Println(v)
      }
      errs = []error{}
    }

    fmt.Println("")

  }
}
