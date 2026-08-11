// hard test 9
// sqlite basic

try {
  cns db_8 = db("sqlite", "file::memory:?cache=shared", "CREATE TABLE test (id INT);")
  pt(get(db_8))
} err(e) {
  pt(e)
}
