// king test 1
// map reduce simulation

cns data_0 = [1, 2, 3, 4, 5]
mapped_0 = []
j_0 = 0
lp len(data_0) {
  mapped_0 = mapped_0 + [data_0[j_0] * 2]
  j_0 = j_0 + 1
}
sum_0 = 0
k_0 = 0
lp len(mapped_0) {
  sum_0 = sum_0 + mapped_0[k_0]
  k_0 = k_0 + 1
}
pt(sum_0)
