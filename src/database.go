package main










type NameDB struct {
  names map [string] BitVect
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
