// king test 14
// api polling simulation

cns attempts_13 = 0
fn poll_13() {
  cns j_13 = 0
  lp 5 {
    attempts_13 = attempts_13 + 1
    if attempts_13 = 3 {
      rev "SUCCESS"
    }
    j_13 = j_13 + 1
  }
  rev "FAIL"
}
pt(poll_13())
