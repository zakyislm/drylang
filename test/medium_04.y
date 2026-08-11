// medium test 4
// loop continue (con)

cns res_3 = 0
cns i_3 = 0
lp 5 {
  i_3 = i_3 + 1
  if i_3 = 2 {
    con
  }
  res_3 = res_3 + i_3
}
pt(res_3)
