// king test 6
// pub sub simulation

subs_5 = []
fn sub_5(cb) {
  subs_5 = subs_5 + [cb]
}
fn pub_5(msg) {
  j_5 = 0
  lp len(subs_5) {
    cns cb = subs_5[j_5]
    cb(msg)
    j_5 = j_5 + 1
  }
}
fn on_msg_5(m) { pt("Received: " + m) }
sub_5(on_msg_5)
pub_5("Hello World")
