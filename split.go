package piscine

func Split(s, sep string) []string {
var r []string
st := 0

  for i:= 0; i <= len(s) - len(sep) ; i ++ {
    if s[i:i + len(sep)] == sep {
      r = append(r, s[st:i])
      st = i + len(sep)
      i += len(sep) - 1
    }
  }
  r = append(r, s[st:])
  return r
}
