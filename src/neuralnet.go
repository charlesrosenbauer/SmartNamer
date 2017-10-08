package main



import (
  "math/rand"
  "strings"
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










func formatConcat (ss []string, c Case, l Capitalization) string {
  ss[0] = strings.ToLower(ss[0])
  for i := 1; i < len(ss); i++ {
    if (l == UPPERCASE) || (c == CAMELCASE) {
      ss[i] = strings.ToUpper(ss[i])
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
