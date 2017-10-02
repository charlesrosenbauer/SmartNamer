package main



import "fmt"










type NameDB struct {
  names map [string] BitVect
}










func (db NameDB) showDB () {
  for k, v := range db.names {
    fmt.Printf("%s : %x %x %x %x %x %x %x %x \n",
      k, v.bits[0], v.bits[1], v.bits[2], v.bits[3], v.bits[4], v.bits[5], v.bits[6], v.bits[7])
  }
}










func (db NameDB) addFile (fname string, ids []StringPos) {

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










func (db NameDB) extractArray () (ret []StringRepPair) {

  ret = make([]StringRepPair, len(db.names))

  i := 0
  for k, v := range db.names {
    ret[i] = StringRepPair{k, v}
    i++
  }

  return ret
}
