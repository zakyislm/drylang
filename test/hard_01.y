// hard test 1
// try err inside loop

cns errors_0 = 0
lp 3 {
  try {
    pt(1 / 0)
  } err(e) {
    errors_0 = errors_0 + 1
  }
}
pt(errors_0)
