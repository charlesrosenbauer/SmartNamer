package main



import (
  "io/ioutil"
  //"regexp"
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










/*
  Hardcoded map literal for getting a regular expression from a file extension.
  The purpose of this is to handle extracting identifiers from languages with different
    requirements for identifiers.
  A more robust solution, which will hopefully be added later, is to have these loaded
    from an XML or something. That way more options can be added later.
*/
var extMap = map [string] string {
  // C/C++
  ".c"    : "[a-zA-Z\\d_]+",
  ".h"    : "[a-zA-Z\\d_]+",
  ".cc"   : "[a-zA-Z\\d_]+",
  ".hh"   : "[a-zA-Z\\d_]+",
  ".cpp"  : "[a-zA-Z\\d_]+",
  ".hpp"  : "[a-zA-Z\\d_]+",

  //Haskell
  ".hs"   : "[a-zA-Z\\d]+",
  ".lhs"  : "[a-zA-Z\\d]+",

  //My personal language Bzo, because why the hell not?
  ".bz"   : "[^ \\$ \\( \\) \\] \\[ \\{ \\}: ; \\. \\, \\_ \\` \\\" \\@ ]+",
  ".lbz"  : "[^ \\$ \\( \\) \\] \\[ \\{ \\}: ; \\. \\, \\_ \\` \\\" \\@ ]+",

  //Jai, because why the hell not?
  ".jai"  : "[a-zA-Z\\d_]+",

  //Python
  ".py"   : "[a-zA-Z\\d_]+",
  ".pyc"  : "[a-zA-Z\\d_]+",
  ".pyd"  : "[a-zA-Z\\d_]+",
  ".pyo"  : "[a-zA-Z\\d_]+",
  ".pyw"  : "[a-zA-Z\\d_]+",
  ".pyz"  : "[a-zA-Z\\d_]+",

  //Javascript
  ".js"   : "[a-zA-Z\\d_]+",

  //Golang
  ".go"   : "[a-zA-Z\\d_]+",

  //Rust
  ".rs"   : "[a-zA-Z\\d_]+",
  ".rlib" : "[a-zA-Z\\d_]+",

  //C#
  ".cs"   : "[a-zA-Z\\d_]+",

  //Elm
  ".elm"  : "[a-zA-Z\\d_]+",

  //Clojure
  ".clj"  : "[^ \\( \\) \\] \\[ \\} \\{ \\;]+",
  ".cljs" : "[^ \\( \\) \\] \\[ \\} \\{ \\;]+",
  ".cljc" : "[^ \\( \\) \\] \\[ \\} \\{ \\;]+",
  ".edn"  : "[^ \\( \\) \\] \\[ \\} \\{ \\;]+",

  //Kotlin
  ".kt"  : "[a-zA-Z\\d_]+",
  ".kts" : "[a-zA-Z\\d_]+",

  //Idris
  ".idr" : "[a-zA-Z\\d_]+",
  ".lidr": "[a-zA-Z\\d_]+",

  //Erlang
  ".erl" : "[a-zA-Z\\d_]+",
  ".hlr" : "[a-zA-Z\\d_]+",

  //Scala
  ".scala": "[a-zA-Z\\d_]+",
  ".sc"   : "[a-zA-Z\\d_]+",

  //ML / Caml / OCaml
  ".ml"   : "[a-zA-Z\\d_]+",
  ".mli"  : "[a-zA-Z\\d_]+",
}









/* !!HARD HAT ZONE!!
func getIds (reg, text, fname string) ([]string, error) {
  failArr := []string{}

  regex, err0 := regexp.Compile(reg)
  if err0 == nil {
    return failArr, err0
  }

  langRegex := regexp.MustCompile("\\.[a-zA-Z]+")
  lang := langRegex.FindString(fname)  //get file extension to figure out the language


}
*/
