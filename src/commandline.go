package main



import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "errors"
  "sort"
)










func printHelpScreen() {
  printStartScreen()
  fmt.Println("\nCommands: ")
  fmt.Println("add-to-workspace [fs] : Add source file to current workspace. Does not actually load the file or perform safety checks.")
  fmt.Println("show-workspace        : Show all files currently in the workspace. They are not necessarily loaded, however.")
  fmt.Println("clear-workspace       : Remove all files from current workspace.")
  fmt.Println("load-files       [fs] : Loads contents of any files in the workspace into memory. Optionally can add/load additional files in the process.")
  fmt.Println("fold                  : Extracts semantic and contextual info from all loaded files via semantic folding.")
  fmt.Println("show-fold             : Shows data from semantic folding. This is for experts and debugging.")
  fmt.Println("show-word-fold        : Similar to show-fold, but shows data from the word database.")
  fmt.Println("add-words        [ws] : Adds words to the word database. These are combined for Id suggestions.")
  fmt.Println("query-words      [ws] : Provide a set of words to determine which ones are in the word database.")
  fmt.Println("show-words            : Shows all words in the word database (a lot of text).")
  fmt.Println("similar-words    [ws] : Shows top 5 words in the word database that appear similar to provided words. Provides unusual results on occasion.")
  fmt.Println("get-case              : Get letter case used for ID suggestions.")
  fmt.Println("set-case       [case] : Set letter case used for ID suggestions. Options are camelcase, snakecase, and kebabcase.")
  fmt.Println("set-upper             : Set all ID suggestions to start with an uppercase letter.")
  fmt.Println("set-lower             : Set all ID suggestions to start with a lowercase letter.")
  fmt.Println("get-capitalization    : Check if ID suggestions are set to start with uppercase or lowercase letters.")
  //fmt.Println("learn                 : Trains neural network for ID prediction. May take a while.")
  //fmt.Println("relearn               : Resets neural network, then runs implicit learn command.")
  fmt.Println("predict          [ws] : Takes a list of IDs and provides suggestions for better replacements.")
  fmt.Println("?                     : Show Help Screen (You Are Here).")
  fmt.Println("quit                  : Quit the program.")

  fmt.Println("\n")
}










func printStartScreen() {
    fmt.Println("\nSmart Namer, created by Charles Rosenbauer")
    fmt.Println("https://github.com/charlesrosenbauer/SmartNamer for info.")
    fmt.Println("Enter \"?\" for help.\n")
}










type Case int
const (
  CAMELCASE Case = iota
  SNAKECASE Case = iota
  KEBABCASE Case = iota
)










type Capitalization int
const (
  UPPERCASE Capitalization = iota
  LOWERCASE Capitalization = iota
)










func commandLoop() {

  // System State
  var (
    errs  []error
  )
  files := map[string]string{}
  dbpar := map[string]BitVect{}
  db := NameDB{dbpar}
  worddb := NameDB{map[string]BitVect{}}
  for _, v := range words {
    worddb.names[v] = representID(v)
  }

  lettercase := CAMELCASE
  capitlcase := LOWERCASE

  var predictor Predictor
  predictor.New()

  // Command Loop State
  reader := bufio.NewReader(os.Stdin)
  cont := true

  for cont {
    fmt.Print(">>> ")
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println(err)
    }
    commandtemp := strings.Split(strings.Trim(text, "\n"), " ")
    var command []string
    for _, v := range commandtemp {
      if v != "" {
        command = append(command, v)
      }
    }

    if len(command) > 0 {
      switch command[0] {



      case "quit" :
        cont = false



      case "get-case" : {
        switch lettercase {

          case CAMELCASE :
            fmt.Println("Camel case (camelCase)")

          case SNAKECASE :
            fmt.Println("Snake case (snake_case)")

          case KEBABCASE :
            fmt.Println("Kebab case (kebab-case)")

          default        : errs = append(errs, errors.New("Unexpected Letter case Value. Alert the developers!"))
        }
      }



      case "set-case" : {
        faulterr := errors.New("No suitable case value provided.\n    Valid values are \"camelcase\", \"snakecase\", and \"kebabcase\"")
        if len(command) > 1 {
          switch command[1] {

            case "camelcase" :
              lettercase = CAMELCASE

            case "snakecase" :
              lettercase = SNAKECASE

            case "kebabcase" :
              lettercase = KEBABCASE

            default :
              errs = append(errs, faulterr)
          }
        }else{
          errs = append(errs, faulterr)
        }
      }



      case "set-upper" : {
        capitlcase = UPPERCASE
      }



      case "set-lower" : {
        capitlcase = LOWERCASE
      }



      case "get-capitalization" : {
        switch capitlcase {

        case UPPERCASE:
          fmt.Println("Uppercase")

        case LOWERCASE:
          fmt.Println("Lowercase")

        default:
          errs = append(errs, errors.New("Unexpected Capitalization case value. Alert the developers!"))
        }
      }



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



      case "fold" : {
        var ids [][]StringPos
        var fnames []string


        // Reset Databases
        db = NameDB{map[string]BitVect{}}
        for _, v := range words {
          worddb.names[v] = representID(v)
        }

        for k, v := range files {
          fnames = append(fnames, k)
          idsTemp, err := getIds(v, k)

          if err != nil {
            errs = append(errs, err)
          }else{
            ids = append(ids, idsTemp)
          }
        }

        for i, v := range ids {
          db.addFile(fnames[i], v)
        }

        // Load any applicable contexts into worddb
        tempdb := worddb
        for i, val := range db.names {
          key := strings.ToLower(i)
          _, ok := worddb.names[key]
          if ok {
            tempdb.names[key] = val
          }
        }
        worddb = tempdb
      }



      case "show-fold" : {
        db.showDB()
      }



      case "show-word-fold" : {
        worddb.showDB()
      }



      case "add-words" : {
        for i := 1; i < len(command); i++ {
          worddb.names[command[i]] = representID(command[i])
        }
      }



      case "query-words" : {
        for i := 1; i < len(command); i++ {
          _, ok := worddb.names[command[i]]
          if ok {
            fmt.Println(command[i], " is recorded")
          }else{
            fmt.Println(command[i], " is not recorded")
          }
        }
      }



      case "similar-words" : {
        for i := 1; i < len(command); i++ {
          word := representID(command[i])
          val, ok := db.names[command[i]]
          if ok {
            word = val
          }

          list := worddb.findSimilar(word, 5)
          fmt.Println("Words similar to: ", command[i], ": ")
          for _, v := range list {
            fmt.Println("    ", v)
          }
          fmt.Println("\n")
        }
      }



      case "show-words" : {
        var list []string
        for i, _ := range worddb.names {
          list = append(list, i)
        }
        sort.Strings(list)
        for _, v := range list {
          fmt.Println(v)
        }
      }


/*
      case "learn" : {

      }



      case "relearn" : {

      }
*/


      case "predict" : {
        for i := 1; i < len(command); i++ {
          word, ok := db.names[command[i]]
          if ok {
            list := predictor.predictWords(word, 5, &worddb)
            fmt.Println("Prediction: ", formatConcat(list, lettercase, capitlcase))
          }else{
            fmt.Println("Identifier", command[i], "does not exist in the code.")
          }
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
