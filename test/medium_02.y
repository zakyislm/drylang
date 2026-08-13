// medium test 2
// nested loop

cns c_1 = 0
lp 3 {
  lp 3 {
    c_1 = c_1 + 1
  }
}
pt(c_1)
