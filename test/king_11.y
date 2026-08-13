// king test 11
// error stacking

fn fail_10() {
  err "deep error"
}
fn wrap_10() {
  try {
    fail_10()
  } err(e) {
    err "wrapped: " + e
  }
}
try {
  wrap_10()
} err(e) {
  pt(t)
}
