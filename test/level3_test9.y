
    val_aicJtV = 42 % 3
    on (val_aicJtV) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_aicJtV)
    

    sum_tMcsuO = 0
    lp 15 {
        sum_tMcsuO = sum_tMcsuO + 1
        if (sum_tMcsuO > 100) { done }
    }
    pt(sum_tMcsuO)
    

    fn calc_gqQOaY(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_rqBXHN = calc_gqQOaY(73, 87+1)
    pt(res_rqBXHN)
    

    val_XhEtyQ = 2 % 3
    on (val_XhEtyQ) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_XhEtyQ)
    

    use "dummy.y"
    
    e_WjSnIN = enc.b64("hello drylang")
    pt(e_WjSnIN)
    
    j_ayXkJV = json(`{"test": 123}`)
    pt(j_ayXkJV)
    
    // Test get() for type info
    val_UVTajz = get("hello")
    pt(val_UVTajz)
    
