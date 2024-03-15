package piscine

func Max(a []int) int {
  max := a[0]
if len(a) <= 0 {
    return 0
  }
for _,n := range a {
if n > max {
  max = n
}
}
return max
}
