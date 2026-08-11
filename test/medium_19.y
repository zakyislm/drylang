// medium test 19
// file read

try {
  cns content_18 = read("testres.log")
  pt(len(content_18) > 0)
} err(e) {
  pt(t)
}
