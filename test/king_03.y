// king test 3
// awt inside loop

sum_2 = 0
asn fn fetch_2(id) {
  sum_2 = sum_2 + (id * 10)
}
asn fn main_2() {
  j_2 = 1
  lp 5 {
    uni fetch_2(j_2)
    j_2 = j_2 + 1
  }
  awt
  pt(sum_2)
}
main_2()
