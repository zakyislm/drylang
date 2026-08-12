// hard test 11
// write and read file

try {
  w("temp_10.txt", "data")
  cns content_10 = r("temp_10.txt")
  pt(content_10 = "data")
  rm("temp_10.txt")
} err(e) {
  pt(e)
}
