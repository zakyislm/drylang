// medium test 8
// closure

fn make_adder_7(x) {
  fn inner_7(y) {
    rev x + y
  }
  rev inner_7
}
cns add5_7 = make_adder_7(5)
pt(add5_7(10))
