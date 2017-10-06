package main










const THRESHOLD = 0.9










type NetLayer struct {
  weights [256][512]float32
}










func (net *NetLayer) predict (word BitVect) BitVect {
  var input [512]float32
  for i := 0; i < 512; i++ {
    var x uint = uint(i) % 64
    var y uint = uint(i) / 64
    input[i] = 0.0;
    if (word.bits[y] & (1 << x)) != 0 {
      input[i] = 1.0;
    }
  }

  var ret BitVect
  var accum float32 = 0
  for n := 0; n < 256; n++ {
    for w := 0; w < 512; w++ {
      accum += net.weights[n][w] * input[w]
    }
    if accum > THRESHOLD {
      var x uint = uint(n) % 64
      var y uint = uint(n) / 64
      ret.bits[y] |= (1 << x)
    }
  }

  return ret

}
