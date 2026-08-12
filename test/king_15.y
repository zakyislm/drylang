// king test 15
// concurrent math

cns async_sum_14 = 0
asn fn calc_14(x) {
  async_sum_14 = async_sum_14 + (x * x)
}
asn fn main_14() {
  uni calc_14(10)
  uni calc_14(20)
  awt
  pt(async_sum_14)
}
main_14()
