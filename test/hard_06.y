// hard test 6
// mul concurrent call

cns res_5 = 0
asn fn job_5() {
  res_5 = res_5 + 1
}
mul 2 job_5()
awt
pt(res_5 > 0)
