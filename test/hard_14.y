// hard test 14
// higher order functions

fn map_func_13(fn_arg, val) {
  rev fn_arg(val)
}
fn double_13(x) {
  rev x * 2
}
pt(map_func_13(double_13, 5))
