// medium test 5
// recursive function

fn fib_4(n) {
  if n < 2 {
    rev n
  }
  rev fib_4(n-1) + fib_4(n-2)
}
pt(fib_4(6))
