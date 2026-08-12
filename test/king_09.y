// king test 9
// bubble sort

arr_8 = [5, 2, 9, 1, 5, 6]
j_8 = 0
lp len(arr_8) {
  k_8 = 0
  lp len(arr_8) - 1 {
    if arr_8[k_8] > arr_8[k_8+1] {
      tmp = arr_8[k_8]
      arr_8[k_8] = arr_8[k_8+1]
      arr_8[k_8+1] = tmp
    }
    k_8 = k_8 + 1
  }
  j_8 = j_8 + 1
}
pt(arr_8[0])
