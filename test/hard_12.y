// hard test 12
// deep recursion

fn recur_11(n) {
  if n = 0 {
    rev 0
  }
  rev 1 + recur_11(n - 1)
}
pt(recur_11(20))
