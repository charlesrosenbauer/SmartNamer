package main






type BitVect struct {
  bits [8]uint64
}










func fixTo27 (char byte) uint {
  var (
    a uint = 'a'
    z uint = 'z'
    A uint = 'A'
    Z uint = 'Z'
    c uint = uint(char)
  )
  if ((c >= a) && (c <= z)) {
    return c - a
  }else if ((c >= A) && (c <= Z)){
    return c - A
  }else{
    return 26
  }
}










func insertIdBit (x, y, z uint, v BitVect) BitVect {
  var pos uint = (x * 729) + (y * 27) + z
  var a   uint = pos / 64
  var b   uint = pos % 64
  v.bits[a] |= (1 >> b)
  return v
}










func representID (id string) BitVect {
  var vect BitVect

  length := len(id)
  if (length == 1){
    vect = insertIdBit(fixTo27(id[0]), 26, 26, vect)
  }else if (length == 2){
    x   := fixTo27(id[0])
    y   := fixTo27(id[1])
    vect = insertIdBit(x, y , 26, vect)
    vect = insertIdBit(y, 26, 26, vect)
  }else{
    for i := 0; i < length-2; i++ {
      x   := fixTo27(id[i])
      y   := fixTo27(id[i+1])
      z   := fixTo27(id[i+2])
      vect = insertIdBit(x, y, z, vect)
    }
    x   := fixTo27(id[length-2])
    y   := fixTo27(id[length-1])
    vect = insertIdBit(x, y , 26, vect)
    vect = insertIdBit(y, 26, 26, vect)
  }
  return vect
}
