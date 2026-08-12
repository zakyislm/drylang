// hard test 10
// regex initialization

try {
  cns rgx_9 = reg("[a-z]+")
  pt(get(rgx_9))
} err(e) {
  pt(t)
}
