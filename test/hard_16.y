// hard test 16
// logical AND shortcut

fn side_effect_15() {
  pt("called")
  rev t
}
pt(f & side_effect_15())
