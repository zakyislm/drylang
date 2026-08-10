// test_04_loops.y

// Counted loop
sum = 0
lp 5 {
    sum = sum + 1
}
pt sum

// Infinite loop with break and continue
i = 0
lp {
    i = i + 1
    if i = 3 {
        con
    }
    if i > 5 {
        done
    }
    pt i
}

// Top level variable shadowing fix test
g_var = 0
lp 3 {
    g_var = g_var + 1
}
pt g_var // Should be 3
