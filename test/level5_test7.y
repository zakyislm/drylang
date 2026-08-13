
    fn calc_Ilqyun(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_DYnwni = calc_Ilqyun(15, 16+1)
    pt(res_DYnwni)
    

    val_mQWkub = 86 % 3
    on (val_mQWkub) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_mQWkub)
    

    use "dummy.y"
    
    e_KjBLkX = enc.b64("hello drylang")
    pt(e_KjBLkX)
    
    j_ofDWJd = json(`{"test": 123}`)
    pt(j_ofDWJd)
    
    // Test get() for type info
    val_GojNDN = get("hello")
    pt(val_GojNDN)
    

    cns C_uEZaoL = 29
    v_wxESzw = unknown
    v_wxESzw = C_uEZaoL + 13
    if (v_wxESzw > 0) {
        v_wxESzw = v_wxESzw * 2
    } el {
        v_wxESzw = 0
    }
    pt(v_wxESzw)
    

    val_ZyAekZ = 18 % 3
    on (val_ZyAekZ) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_ZyAekZ)
    

    use "dummy.y"
    
    e_gPUhGG = enc.b64("hello drylang")
    pt(e_gPUhGG)
    
    j_kLBNwz = json(`{"test": 123}`)
    pt(j_kLBNwz)
    
    // Test get() for type info
    val_DolTSA = get("hello")
    pt(val_DolTSA)
    

    sum_ikbEHg = 0
    lp 15 {
        sum_ikbEHg = sum_ikbEHg + 1
        if (sum_ikbEHg > 100) { done }
    }
    pt(sum_ikbEHg)
    
