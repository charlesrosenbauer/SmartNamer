package main



import (
  "fmt"
  "sort"
)










type NameDB struct {
  names map [string] BitVect
}










func (db NameDB) showDB () {
  for k, v := range db.names {
    fmt.Printf("%s : %x %x %x %x %x %x %x %x \n",
      k, v.bits[0], v.bits[1], v.bits[2], v.bits[3], v.bits[4], v.bits[5], v.bits[6], v.bits[7])
  }
}










func (db *NameDB) addFile (fname string, ids []StringPos) {

  for i:=0; i<len(ids); i++ {
    id := ids[i].str
    _, in := db.names[id]
    if !in {
      x := representID(ids[i].str)
      x  = representPosition(fname, ids[i].pos, x)
      db.names[id] = x
    }else{
      db.names[id] = representPosition(fname, ids[i].pos, db.names[id])
    }
  }
}










type StringRepPair struct {
  str string
  rep BitVect
}










func (db *NameDB) extractArray () (ret []StringRepPair) {

  ret = make([]StringRepPair, len(db.names))

  i := 0
  for k, v := range db.names {
    ret[i] = StringRepPair{k, v}
    i++
  }

  return ret
}










// Because I can't make interfaces with anonymous structs apparently
type StringFloatList struct{
  x []struct{
    s string
    f float32
  }
}










func (s StringFloatList) Len () int {
  return len(s.x)
}










func (s StringFloatList) Less (i, j int) bool {
  return s.x[i].f < s.x[j].f
}










func (s StringFloatList) Swap (i, j int) {
  temp := s.x[i]
  s.x[i] = s.x[j]
  s.x[j] = temp
}










func (db *NameDB) findSimilar (comp BitVect, num int) []string {
  var list []struct{s string; f float32}

  for i, v := range db.names {
    unionPop := float32(vectPopulation(vectUnion(v, comp)))
    if unionPop < 1 {
      unionPop = 0.001  // stop x/0 errors
    }
    interPop := float32(vectMatch(v, comp))
    var div float32 = interPop / unionPop
    item := struct{s string; f float32}{i, div}
    list = append(list, item)
  }
  newlist := StringFloatList{list}
  sort.Sort(newlist)

  var ret []string
  for i := 1; i <= num; i++ {
    index := len(newlist.x) - i
    if index >= 0 {
      ret = append(ret, newlist.x[index].s)
    }
  }
  return ret
}
