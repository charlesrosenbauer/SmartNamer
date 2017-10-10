package main



import (
  "math/rand"
  "strings"
  "math"
)










const THRESHOLD = 0.9










type NetLayer struct {
  weights [256][512]float32
}










type Predictor struct {
  standardLayers [5]NetLayer
  contextLayers  [4]NetLayer
}










func (net *NetLayer) New () {
  for i, v := range net.weights {
    for j, _ := range v {
      net.weights[i][j] = (2.0 * rand.Float32()) - 1.0
    }
  }
}










func (pred *Predictor) New () {
  for i, _ := range pred.standardLayers {
    pred.standardLayers[i].New()
  }
  for i, _ := range pred.contextLayers {
    pred.contextLayers[i].New()
  }
}










func (net *NetLayer) predict (word BitVect) (BitVect, *[256]float32) {
  var input [512]float32
  for i := 0; i < 512; i++ {
    var x uint = uint(i) % 64
    var y uint = uint(i) / 64
    input[i] = 0.0
    if (word.bits[y] & (1 << x)) != 0 {
      input[i] = 1.0
    }
  }

  var ret BitVect
  var accums [256]float32
  for n := 0; n < 256; n++ {
    accums[n] = 0
    for w := 0; w < 512; w++ {
      accums[n] += net.weights[n][w] * input[w]
    }
    if accums[n] > THRESHOLD {
      var x uint = uint(n) % 64
      var y uint = uint(n) / 64
      ret.bits[y] |= (1 << x)
    }
  }

  return ret, &accums
}










func (net *NetLayer) learn (in, out BitVect, rate float32, accums *[256]float32) {
  var input  [512]float32
  var output [256]float32
  for i := 0; i < 512; i++ {
    var x uint = uint(i) % 64
    var y uint = uint(i) / 64
    input[i] = 0.0
    if (in.bits[y] & (1 << x)) != 0 {
      input[i] = 1.0
    }
    if (out.bits[y] & (1 << x)) != 0 {
      output[i] = 1.0
    }
  }

  for n := 0; n < 256; n++ {
    var delta float32 = rate * (output[n] - accums[n])
    for w := 0; w < 512; w++ {
      if input[w] > 0.0 {
        net.weights[n][w] += net.weights[n][w] * delta
      }
    }
  }
}










func (pred *Predictor) predictWords (in BitVect, numwords int, worddb *NameDB) [][]string {

  maxwords := 5
  if (numwords > 0) && (numwords < 5){
    maxwords = numwords
  }

  var stdVects [5]BitVect
  for i := 0; i < maxwords; i++ {
    stdVects[i], _ = pred.standardLayers[i].predict(in)
  }

  var ctxVects [4]BitVect
  for i := 0; i < maxwords-1; i++ {
    ctxVects[i], _ = pred.contextLayers[i].predict(stdVects[i])
  }

  var ret [][]string
  ret = append(ret, (worddb.findSimilar(stdVects[0], 5)))
  for i := 1; i < maxwords; i++ {
    stdVects[i] = vectUnion(stdVects[i], ctxVects[i-1])
    ret = append(ret, (worddb.findSimilar(stdVects[i], 5)))
  }

  return ret
}










func (pred *Predictor) learnWord (in, out BitVect, c Case, l Capitalization, worddb *NameDB, numwords int, rate float32) float32 {

  // Get Outputs of Predictor

  if numwords < 1 {
    numwords = 1
  }else if numwords > 5 {
    numwords = 5
  }

  var stdVects [5]BitVect
  var stdAccum [5]*[256]float32
  for i := 0; i < numwords; i++ {
    stdVects[i], stdAccum[i] = pred.standardLayers[i].predict(in)
  }

  var ctxVects [4]BitVect
  var ctxAccum [4]*[256]float32
  for i := 0; i < numwords-1; i++ {
    ctxVects[i], ctxAccum[i] = pred.contextLayers[i].predict(stdVects[i])
  }


  // Get some string predictions

  var attempts [][]string
  attempts = append(attempts, (worddb.findSimilar(stdVects[0], 5)))
  for i := 1; i < numwords; i++ {
    vects := vectUnion(stdVects[i], ctxVects[i-1])
    attempts = append(attempts, (worddb.findSimilar(vects, 5)))
  }

  _, veclist := randomStringBitPredictions(attempts, c, l, 20)

  var minvec []BitVect
  minflt := float32(math.MaxFloat32)
  for _, v := range veclist {
    tmpvect := reduceVect(v, BitVect{[8]uint64{0, 0, 0, 0, 0, 0, 0, 0}}, vectUnion)
    tmpflt  := measureSimilarity(tmpvect, out)
    if tmpflt < minflt {
      minvec = v
      minflt = tmpflt
    }
  }


  // Learn based on minvec

  for i, v := range minvec {
    if i == 0 {
      stdv := vectIntersection(v, stdVects[i])
      pred.standardLayers[0].learn(in, stdv, rate, stdAccum[0])
    }else{
      stdv := vectIntersection(v, stdVects[i])
      pred.standardLayers[i].learn(in, stdv, rate, stdAccum[i])
      ctxv := vectIntersection(v, ctxVects[i-1])
      pred.standardLayers[i].learn(stdVects[i-1], ctxv, rate, ctxAccum[i-1])
    }
  }

  return minflt
}










func formatConcat (ss []string, c Case, l Capitalization) string {
  ss[0] = strings.ToLower(ss[0])
  for i := 1; i < len(ss); i++ {
    if (c == CAMELCASE) || (l == UPPERCASE) {
      ss[i] = strings.Title(ss[i])
    }else{
      ss[i] = strings.ToLower(ss[i])
    }
  }
  switch c {
    case CAMELCASE :
      return strings.Join(ss, "")

    case SNAKECASE :
      return strings.Join(ss, "_")

    case KEBABCASE :
      return strings.Join(ss, "-")
  }
  return ""
}










/*
  Note: ss is a collection of collections of strings;
    Each output string will consist of N substrings concatenated (with some additional transforms).
    ss consists of N collections of substrings.
    Each of said collections are selected from randomly to provide a random substring.

    For example, given
      ss := [[a b c d]
             [e f g]
             [h i j k l]]
      , examples of possible outputs are "afi", "cgl", and "dej", not taking into account formatting transforms.
*/
func randomStringPredictions(ss [][]string, c Case, l Capitalization, n int) []string {
  ret := make([]string, n)
  for i, _ := range ret {
    // For each ID to be predicted

    var outputstr []string
    for _, strs := range ss {
      // For each collection of collections

      size := len(strs)
      outputstr = append(outputstr, strs[rand.Int() % size])
    }

    ret[i] = formatConcat(outputstr, c, l)
  }
  return ret
}










func randomStringBitPredictions(ss [][]string, c Case, l Capitalization, n int) ([]string, [][]BitVect) {
  retStr := make([]string, n)
  retVec := make([][]BitVect, n)
  for i, _ := range retStr {
    // For each ID to be predicted

    var outputstr []string
    var outputvec []BitVect
    for _, strs := range ss {
      // For each collection of collections

      size := len(strs)
      str  := strs[rand.Int() % size]
      outputstr = append(outputstr, str)
      outputvec = append(outputvec, representID(str))
    }

    retStr[i] = formatConcat(outputstr, c, l)
    retVec[i] = outputvec
  }
  return retStr, retVec
}










func errorMetric (unionPop, interPop float32) float32 {
  if interPop == 0 {
    interPop = 0.01
  }
  return unionPop / interPop
}
