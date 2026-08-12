// hard test 13
// nested try err

try {
  try {
    pt(1 / 0)
  } err(e1) {
    pt("inner")
    pt(1 / 0)
  }
} err(e2) {
  pt("outer")
}
