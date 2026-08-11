// medium test 20
// http req basic

try {
  cns resp_19 = req("GET", "https://httpbin.org/get")
  pt(get(resp_19))
} err(e) {
  pt(e)
}
