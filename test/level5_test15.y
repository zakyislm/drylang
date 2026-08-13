
    sum_roHiLD = 0
    lp 8 {
        sum_roHiLD = sum_roHiLD + 1
        if (sum_roHiLD > 100) { done }
    }
    pt(sum_roHiLD)
    

    fn calc_LzLXPj(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ufLrcC = calc_LzLXPj(64, 28+1)
    pt(res_ufLrcC)
    

    sum_BSCKpg = 0
    lp 14 {
        sum_BSCKpg = sum_BSCKpg + 1
        if (sum_BSCKpg > 100) { done }
    }
    pt(sum_BSCKpg)
    

    try {
        throw_vZXRkx = unknown
        throw_vZXRkx()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_NfAjyn = enc.b64("hello drylang")
    pt(e_NfAjyn)
    
    j_RZjeAw = json(`{"test": 123}`)
    pt(j_RZjeAw)
    
    // Test get() for type info
    val_SiSHLA = get("hello")
    pt(val_SiSHLA)
    

    use "dummy.y"
    
    e_pYbZfd = enc.b64("hello drylang")
    pt(e_pYbZfd)
    
    j_CbSAkO = json(`{"test": 123}`)
    pt(j_CbSAkO)
    
    // Test get() for type info
    val_ajSAnA = get("hello")
    pt(val_ajSAnA)
    

    val_gfQhQt = 2 % 3
    on (val_gfQhQt) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_gfQhQt)
    
