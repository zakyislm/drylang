
    sum_hhjXVi = 0
    lp 6 {
        sum_hhjXVi = sum_hhjXVi + 1
        if (sum_hhjXVi > 100) { done }
    }
    pt(sum_hhjXVi)
    

    arr_pcAaFk = [78, 42+1, 16+2]
    map_hMxgqZ = {"a": arr_pcAaFk[0], "b": arr_pcAaFk[1]}
    pt(map_hMxgqZ["a"])
    

    val_UKOdBi = 86 % 3
    on (val_UKOdBi) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_UKOdBi)
    

    asn fn async_task_qSOxEy(x) {
        rev x * 2
    }
    uni async_task_qSOxEy(80)
    
    asn fn worker_Ejadqe(y) {
        pt("working on", y)
    }
    mul 2 worker_Ejadqe(52)
    
    awt
    

    try {
        throw_gbTrQi = unknown
        throw_gbTrQi()
    } err(e) {
        pt("caught error")
    }
    

    val_YkHRqG = 49 % 3
    on (val_YkHRqG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_YkHRqG)
    

    use "dummy.y"
    
    e_VUoHoS = enc.b64("hello drylang")
    pt(e_VUoHoS)
    
    j_MzBGUV = json(`{"test": 123}`)
    pt(j_MzBGUV)
    
    // Test get() for type info
    val_MJqVpJ = get("hello")
    pt(val_MJqVpJ)
    
