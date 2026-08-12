// king test 4
// tree traversal recursive

cns tree_3 = {"val": 1, "left": {"val": 2, "left": unknown, "right": unknown}, "right": {"val": 3, "left": unknown, "right": unknown}}
fn sum_tree_3(node) {
  if node = unknown {
    rev 0
  }
  rev node["val"] + sum_tree_3(node["left"]) + sum_tree_3(node["right"])
}
pt(sum_tree_3(tree_3))
