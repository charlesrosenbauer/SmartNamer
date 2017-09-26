package main



//import "math/bits"










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










func vectUnion (a, b BitVect) BitVect {
  for i := 0; i < 8; i++ {
    a.bits[i] |= b.bits[i]
  }
  return a
}










func vectIntersection (a, b BitVect) BitVect {
  for i := 0; i < 8; i++ {
    a.bits[i] &= b.bits[i]
  }
  return a
}










func vectDifference (a, b BitVect) BitVect {
  for i := 0; i < 8; i++ {
    a.bits[i] ^= b.bits[i]
  }
  return a
}










func vectInverse (a BitVect) BitVect {
  for i := 0; i < 8; i++ {
    a.bits[i] = ^(a.bits[i])
  }
  return a
}










func OnesCount64 (x uint64) int {
  /*
    For some reason golang's new math/bits library's not working.
    So screw it. I copied the implementation from online and pasted it here.
    Could probably write a faster version though.
    This implementation isn't particularly efficient.
  */
  const m0 = 0x5555555555555555 // 01010101 ...
  const m1 = 0x3333333333333333 // 00110011 ...
  const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
  const m = 1<<64 - 1
  x = x>>1&(m0&m) + x&(m0&m)
  x = x>>2&(m1&m) + x&(m1&m)
  x = (x>>4 + x) & (m2 & m)
  x += x >> 8
  x += x >> 16
  x += x >> 32
  return int(x) & (1<<7 - 1)
}










func vectPopulation (a, b BitVect) int {
  pcnt := 0
  for i := 0; i < 8; i++ {
    pcnt += OnesCount64(b.bits[i])
  }
  return pcnt
}










func vectMatch (a, b BitVect) int {
  return vectPopulation(vectIntersection(a, b))
}
