package main



import (
  "io/ioutil"
  "regexp"
  "errors"
  "strings"
)










func loadSourceFiles (files []string) ([]string, error) {
  numfiles  := len(files)
  var filetexts []string
  for i := 0; i < numfiles; i++ {
    text, err:= ioutil.ReadFile(files[i])
    if err != nil {
      out := []string{}
      return out, err
    }
    filetexts = append(filetexts, string(text))
  }
  return filetexts, nil
}










func loadSourceFromMap (files map[string]string) (map[string]string, error) {
  ret := map[string]string{}
  for k, _ := range(files) {
    text, err:= ioutil.ReadFile(k)
    if err != nil {
      return ret, err
    }
    ret[k] = string(text)
  }
  return ret, nil
}










/*
  Hardcoded map literal for getting a regular expression from a file extension.
  The purpose of this is to handle extracting identifiers from languages with different
    requirements for identifiers.
  A more robust solution, which will hopefully be added later, is to have these loaded
    from an XML or something. That way more options can be added later.
*/
var extMap = map [string] string {
  // C/C++
  ".c"    : "[a-zA-Z][a-zA-Z\\d_]+",
  ".h"    : "[a-zA-Z][a-zA-Z\\d_]+",
  ".cc"   : "[a-zA-Z][a-zA-Z\\d_]+",
  ".hh"   : "[a-zA-Z][a-zA-Z\\d_]+",
  ".cpp"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".hpp"  : "[a-zA-Z][a-zA-Z\\d_]+",

  //Haskell
  ".hs"   : "[a-zA-Z\\d]+'*",
  ".lhs"  : "[a-zA-Z\\d]+'*",

  //My personal language Bzo, because why the hell not?
  ".bz"   : "[^ $ \\( \\) \\] \\[ \\{ \\} :; \\. \\, \\_ ` \" @ ' \\s]+'*",
  ".lbz"  : "[^ $ \\( \\) \\] \\[ \\{ \\} :; \\. \\, \\_ ` \" @ ' \\s]+'*",

  //Jai, because why the hell not?
  ".jai"  : "[a-zA-Z][a-zA-Z\\d_]+",

  //Python
  ".py"   : "[a-zA-Z][a-zA-Z\\d_]+",
  ".pyc"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".pyd"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".pyo"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".pyw"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".pyz"  : "[a-zA-Z][a-zA-Z\\d_]+",

  //Javascript
  ".js"   : "[a-zA-Z][a-zA-Z\\d_]+",

  //Golang
  ".go"   : "[a-zA-Z][a-zA-Z\\d_]+",

  //Rust
  ".rs"   : "[a-zA-Z][a-zA-Z\\d_]+",
  ".rlib" : "[a-zA-Z][a-zA-Z\\d_]+",

  //C#
  ".cs"   : "[a-zA-Z][a-zA-Z\\d_]+",

  //Elm
  ".elm"  : "[a-zA-Z][a-zA-Z\\d_]+",

  //Clojure
  ".clj"  : "[^ \\( \\) \\] \\[ \\} \\{ ; \\s]+",
  ".cljs" : "[^ \\( \\) \\] \\[ \\} \\{ ; \\s]+",
  ".cljc" : "[^ \\( \\) \\] \\[ \\} \\{ ; \\s]+",
  ".edn"  : "[^ \\( \\) \\] \\[ \\} \\{ ; \\s]+",

  //Kotlin
  ".kt"  : "[a-zA-Z][a-zA-Z\\d_]+",
  ".kts" : "[a-zA-Z][a-zA-Z\\d_]+",

  //Idris
  ".idr" : "[a-zA-Z][a-zA-Z\\d_]+",
  ".lidr": "[a-zA-Z][a-zA-Z\\d_]+",

  //Erlang
  ".erl" : "[a-zA-Z][a-zA-Z\\d_]+",
  ".hlr" : "[a-zA-Z][a-zA-Z\\d_]+",

  //Scala
  ".scala": "[a-zA-Z][a-zA-Z\\d_]+",
  ".sc"   : "[a-zA-Z][a-zA-Z\\d_]+",

  //ML / Caml / OCaml
  ".ml"   : "[a-zA-Z][a-zA-Z\\d_]+",
  ".mli"  : "[a-zA-Z][a-zA-Z\\d_]+",
}










type StringPos struct {
  str string
  pos int
}










func getIds (text, fname string) ([]StringPos, error) {
  failArr := []StringPos{}

  langRegex := regexp.MustCompile("\\.[a-zA-Z]+")
  lang := langRegex.FindString(fname)  //get file extension to figure out the language

  reg, ok := extMap[lang]
  if ! ok {
    return failArr, errors.New("Unrecognized File Extension: " + lang)
  }

  regex, err0 := regexp.Compile(reg)
  if err0 != nil {
    return failArr, err0
  }

  textlines := strings.Split(text, "\n")
  rets := []StringPos{}

  for i:=0; i < len(textlines); i++ {
    ids := regex.FindAllString(textlines[i], -1)
    for j:=0; j < len(ids); j++ {
      rets = append(rets, StringPos{ids[j], i})
    }
  }

  return rets, nil

}
