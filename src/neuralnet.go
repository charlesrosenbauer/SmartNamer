package main










type Arr512 struct {
  vals [512]float32
}










type Arr256 struct {
  vals [256]float32
}










type Perceptron256 struct {
  neurons []Arr256
}










type Perceptron512 struct {
  neurons []Arr512
}










func mul256 (a, b Arr256) Arr256 {
  for i := 0; i < 256; i++ {
    a.vals[i] *= b.vals[i]
  }
  return a
}










func mul512 (a, b Arr512) Arr512 {
  for i := 0; i < 512; i++ {
    a.vals[i] *= b.vals[i]
  }
  return a
}










func dot256 (a, b Arr256) float32 {
  var ret float32 = 0.0
  for i := 0; i < 256; i++ {
    ret += a.vals[i] * b.vals[i]
  }
  return ret
}










func dot512 (a, b Arr512) float32 {
  var ret float32 = 0.0
  for i := 0; i < 512; i++ {
    ret += a.vals[i] * b.vals[i]
  }
  return ret
}










func mul256scalar (a Arr256, b float32) Arr256 {
  for i := 0; i < 256; i++ {
    a.vals[i] *= b
  }
  return a
}










func mul512scalar (a Arr512, b float32) Arr512 {
  for i := 0; i < 512; i++ {
    a.vals[i] *= b
  }
  return a
}










func add256 (a, b Arr256) Arr256 {
  for i := 0; i < 256; i++ {
    a.vals[i] += b.vals[i]
  }
  return a
}










func add512 (a, b Arr512) Arr512 {
  for i := 0; i < 512; i++ {
    a.vals[i] += b.vals[i]
  }
  return a
}










func add256scalar (a Arr256, b float32) Arr256 {
  for i := 0; i < 256; i++ {
    a.vals[i] += b
  }
  return a
}










func add512scalar (a Arr512, b float32) Arr512 {
  for i := 0; i < 512; i++ {
    a.vals[i] += b
  }
  return a
}










func sum256 (a Arr256) float32 {
  var ret float32 = 0.0
  for i := 0; i < 256; i++ {
    ret += a.vals[i]
  }
  return ret
}










func sum512 (a Arr512) float32 {
  var ret float32 = 0.0
  for i := 0; i < 512; i++ {
    ret += a.vals[i]
  }
  return ret
}










func mul512x512 (a, b Arr512) Arr512 {
  for i := 0; i < 512; i++ {
    var sum float32 = 0.0
    val := a.vals[i]
    for j := 0; j < 512; j++ {
       sum += val * b.vals[j]
    }
    a.vals[i] = val
  }
  return a
}










func mul512x256 (a Arr512, b Arr256) Arr512 {
  for i := 0; i < 512; i++ {
    var sum float32 = 0.0
    val := a.vals[i]
    for j := 0; j < 256; j++ {
       sum += val * b.vals[j]
    }
    a.vals[i] = val
  }
  return a
}










func mul256x512 (a Arr256, b Arr512) Arr256 {
  for i := 0; i < 256; i++ {
    var sum float32 = 0.0
    val := a.vals[i]
    for j := 0; j < 512; j++ {
       sum += val * b.vals[j]
    }
    a.vals[i] = val
  }
  return a
}










func mul256x256 (a, b Arr256) Arr256 {
  for i := 0; i < 256; i++ {
    var sum float32 = 0.0
    val := a.vals[i]
    for j := 0; j < 256; j++ {
       sum += val * b.vals[j]
    }
    a.vals[i] = val
  }
  return a
}










func toArr256 (x []float32) Arr256 {
  xlen := len(x)
  var ret Arr256
  if xlen > 256 {
    xlen = 256
  }
  for i := 0; i < xlen; i++ {
    ret.vals[i] = x[i]
  }
  for i := xlen; i < 256; i++ {
    ret.vals[i] = 0
  }
  return ret
}










func toArr512 (x []float32) Arr512 {
  xlen := len(x)
  var ret Arr512
  if xlen > 512 {
    xlen = 512
  }
  for i := 0; i < xlen; i++ {
    ret.vals[i] = x[i]
  }
  for i := xlen; i < 512; i++ {
    ret.vals[i] = 0
  }
  return ret
}










func predict256 (p *Perceptron256, x *Arr256) []float32 {
  ret := make([]float32, len(p.neurons))
  for i, v := range p.neurons {
    ret[i] = dot256(v, x)
  }
  return ret
}










func predict512 (p *Perceptron512, x *Arr512) []float32 {
  ret := make([]float32, len(p.neurons))
  for i, v := range p.neurons {
    ret[i] = dot512(v, x)
  }
  return ret
}










func learn256 (p *Perceptron256, in *Arr256, err float32, eta float32) {
  val := err * eta
  for i, v := range p.neurons {
    p.neurons[i] = add256(v, mul256scalar(in, val))
  }
}










func learn512 (p *Perceptron512, in *Arr512, err float32, eta float32) {
  val := err * eta
  for i, v := range p.neurons {
    p.neurons[i] = add512(v, mul512scalar(in, val))
  }
}










type IdPredictor struct {
  /*
    This is the neural network for predicting identifiers. It takes in a single
    bitvect and returns up to 5 bitvects. The network consists of 5 "layers",
    and learns like a perceptron, though it's structure is a bit unorthodox for
    one. Each layer produces a bitvect output. Inputs for each layer consist of
    the overall input and the output of the previous layer. The first layer only
    takes the input bitvect as an input, as there is no previous layer.
  */

  // Layer A
  wordADirect Perceptron512

  // Layer B
  wordBDirect Perceptron512
  wordBSecond Perceptron256

  // Layer C
  wordCDirect Perceptron512
  wordBSecond Perceptron256

  // Layer D
  wordDDirect Perceptron512
  wordBSecond Perceptron256

  // Layer E
  wordEDirect Perceptron512
  wordBSecond Perceptron256
}









/*
// !!HARD HAT ZONE!!
func (pred *IdPredictor) predict (input BitVect, numWords int) [5]BitVect {
  var ret [5]BitVect
  for i := 0; i < 5; i++ {
    for j := 0; j < 8; j++ {
      ret[i].bits[j] = 0
    }
  }


  if numWords < 1 {
    return ret
  }


  var in Arr512
  for i := 0; i < 512; i++ {
    in.vals[i] = 0.0
    x := i / 64
    y := i % 64
    if (input.bits[x] & (1 << y)) != 0 {
      in.vals[i] = 1.0
    }
  }


  if numWords > 1 {
    ret[0] = predict512(pred.wordADirect, in)
  }


  return ret
}
*/
