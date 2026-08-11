// king test 2
// state machine

state_1 = "INIT"
fn transition_1(s) {
  if s = "INIT" { rev "RUNNING" }
  if s = "RUNNING" { rev "DONE" }
  rev "ERROR"
}
state_1 = transition_1(state_1)
state_1 = transition_1(state_1)
pt(state_1)
