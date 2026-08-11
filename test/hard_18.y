// hard test 18
// multiple awt in expr

sum_17 = 0
asn fn a_17() { sum_17 = sum_17 + 1 }
asn fn b_17() { sum_17 = sum_17 + 2 }
asn fn main_17() {
  uni a_17()
  uni b_17()
  awt
  pt(sum_17)
}
main_17()
