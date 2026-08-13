// hard test 5
// async fn and awt

cns res_4 = 0
asn fn my_async_4() {
  res_4 = 42
}
asn fn main_4() {
  uni my_async_4()
  awt
  pt(res_4)
}
main_4()
