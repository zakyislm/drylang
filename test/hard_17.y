// hard test 17
// logical OR shortcut

fn side_effect_16() {
  pt("called")
  rev t
}
pt(t | side_effect_16())
