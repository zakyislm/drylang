
    use "dummy.y"
    
    e_zedOON = enc.b64("hello drylang")
    pt(e_zedOON)
    
    j_bDcfhA = json(`{"test": 123}`)
    pt(j_bDcfhA)
    
    // Test get() for type info
    val_RoBmMf = get("hello")
    pt(val_RoBmMf)
    

    fn calc_qqxwBi(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_UcUrnp = calc_qqxwBi(68, 37+1)
    pt(res_UcUrnp)
    

    try {
        throw_UbgQla = unknown
        throw_UbgQla()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_HkxVHi(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_fjRbdZ = calc_HkxVHi(3, 34+1)
    pt(res_fjRbdZ)
    

    sum_HKaCre = 0
    lp 5 {
        sum_HKaCre = sum_HKaCre + 1
        if (sum_HKaCre > 100) { done }
    }
    pt(sum_HKaCre)
    

    arr_yWKZvv = [81, 79+1, 42+2]
    map_fiXGXl = {"a": arr_yWKZvv[0], "b": arr_yWKZvv[1]}
    pt(map_fiXGXl["a"])
    

    val_VoRicg = 27 % 3
    on (val_VoRicg) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_VoRicg)
    

    m_BmnFvp = math.sqrt(77)
    h_WyApYM = hash.md5("test_ffZtRz")
    log.inf("hash:", h_WyApYM)
    pt(m_BmnFvp)
    
