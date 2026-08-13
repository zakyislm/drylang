// king test 7
// memoized factorial

cache_6 = {}
fn fact_6(n) {
  if has(cache_6, str(n)) {
    rev cache_6[str(n)]
  }
  if n <= 1 {
    rev 1
  }
  cns res = n * fact_6(n - 1)
  cache_6[str(n)] = res
  rev res
}
pt(fact_6(5))
