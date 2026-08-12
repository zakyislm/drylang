// hard test 8
// cmd execution

try {
  out_7 = cmd("cmd", "/c", "echo hello")
  pt(trm(out_7) = "hello")
} err(e) {
  out2_7 = cmd("sh", "-c", "echo hello")
  pt(trm(out2_7) = "hello")
}
