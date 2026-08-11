// medium test 3
// loop break (done)

cns res_2 = 0
cns i_2 = 0
lp 10 {
  if i_2 = 5 {
    done
  }
  res_2 = res_2 + i_2
  i_2 = i_2 + 1
}
pt(res_2)
