// king test 12
// async file operations

asn fn write_log_11(msg) {
  w("log_11.txt", msg)
}
asn fn run_11() {
  uni write_log_11("test log")
  awt
  cns d = r("log_11.txt")
  pt(d = "test log")
  rm("log_11.txt")
}
run_11()
